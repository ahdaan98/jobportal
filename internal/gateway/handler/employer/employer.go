package employer

import (
	"context"
	"net/http"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/employer"
	"github.com/ahdaan67/jobportal/utils/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctx    context.Context
	client pb.EmployerClient
	cfg    config.Config
}

type Role string

const (
	Employer Role = "employer"
)

func NewHandler(client pb.EmployerClient, cfg config.Config) *Handler {
	return &Handler{
		ctx:    context.Background(),
		client: client,
		cfg:    cfg,
	}
}

func (h *Handler) EmployerSignup(c *gin.Context) {
	var emp CreateEmployerReq
	err := c.ShouldBindJSON(&emp)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"request": "Unable to parse request body"},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	employer, err := h.client.CreateEmployer(h.ctx, toPBEmployerReq(&emp))
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to create employer",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toEmployer(employer)
	token, err := token.GenerateToken(emp.Email, res.ID, string(Employer))
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to generate token",
			Errors:  map[string]string{"token": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Employer Signed Up successfully",
		Data: map[string]interface{}{
			"employer": res,
			"token":    token,
		},
	}
	c.JSON(succRes.Code, succRes)
}
