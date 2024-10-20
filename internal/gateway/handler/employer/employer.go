package employer

import (
	"context"
	"net/http"
	"os"

	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/employer"
	"github.com/ahdaan67/jobportal/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	ctx     context.Context
	client  pb.EmployerClient
	cfg     config.Config
	logger  *logrus.Logger
	logFile *os.File
}

type Role string

const (
	Employer Role = "employer"
)

func NewHandler(client pb.EmployerClient, cfg config.Config, logfile string) *Handler {
	logger, logFile := logging.InitLogrusLogger(logfile)
	return &Handler{
		ctx:     context.Background(),
		client:  client,
		cfg:     cfg,
		logger:  logger,
		logFile: logFile,
	}
}

func (h *Handler) EmployerSignup(c *gin.Context) {
	var emp CreateEmployerReq
	if err := c.ShouldBindJSON(&emp); err != nil {
		h.logger.WithError(err).Error("Failed to parse employer signup request")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"request": "Unable to parse request body"},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Received signup request for employer: %s", emp.Email)

	employer, err := h.client.CreateEmployer(h.ctx, toPBEmployerReq(&emp))
	if err != nil {
		h.logger.WithError(err).Error("Failed to create employer via gRPC")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "failed to create employer",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Employer created successfully: %s", employer.Email)

	res := toEmployer(employer)
	token, err := token.GenerateToken(emp.Email, res.ID, string(Employer))
	if err != nil {
		h.logger.WithError(err).Error("Failed to generate token for employer")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Errors:  map[string]string{"token": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Token generated for employer: %s", emp.Email)

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Employer signed up successfully",
		Data: map[string]interface{}{
			"employer": res,
			"token":    token,
		},
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) LoginEmployer(c *gin.Context) {
	var employerLoginReq EmployerLoginReq
	if err := c.ShouldBindJSON(&employerLoginReq); err != nil {
		h.logger.WithError(err).Error("Failed to parse employer login request")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"request": "Unable to parse request body"},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Employer login attempt: %s", employerLoginReq.Email)

	empRes, err := h.client.LoginEmployer(h.ctx, &pb.EmpLoginReq{
		Email:    employerLoginReq.Email,
		Password: employerLoginReq.Password,
	})
	if err != nil {
		h.logger.WithError(err).Error("Failed employer login")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusUnauthorized,
			Message: "Failed to login",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Employer login successful: %s", employerLoginReq.Email)

	res := toEmployer(empRes)
	token, err := token.GenerateToken(res.Email, res.ID, string(Employer))
	if err != nil {
		h.logger.WithError(err).Error("Failed to generate token for employer")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
			Errors:  map[string]string{"token": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Token generated for employer: %s", res.Email)

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Employer profile retrieved successfully",
		Data: map[string]interface{}{
			"employer": res,
			"token":    token,
		},
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) GetEmployers(c *gin.Context) {
	employers, err := h.client.GetEmployers(c.Request.Context(), &emptypb.Empty{})
	if err != nil {
		h.logger.WithError(err).Error("Failed to get employers")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "failed to get employers",
			Errors: map[string]string{
				"error": err.Error(),
			},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.Infof("Successfully fetched employers")

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "successfully fetched employers",
		Data:    employers,
	}
	c.JSON(succRes.Code, succRes)
}