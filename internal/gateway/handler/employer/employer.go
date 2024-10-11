package employer

import (
    "context"
    "log"
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
        log.Printf("Failed to parse employer signup request: %v", err)
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid input data",
            Errors:  map[string]string{"request": "Unable to parse request body"},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    log.Printf("Received signup request for employer: %s", emp.Email)

    employer, err := h.client.CreateEmployer(h.ctx, toPBEmployerReq(&emp))
    if err != nil {
        log.Printf("Failed to create employer via gRPC: %v", err)
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "failed to create employer",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    log.Printf("Employer created successfully: %s", employer.Email)

    res := toEmployer(employer)
    token, err := token.GenerateToken(emp.Email, res.ID, string(Employer))
    if err != nil {
        log.Printf("Failed to generate token for employer: %s, error: %v", emp.Email, err)
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "failed to generate token",
            Errors:  map[string]string{"token": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    log.Printf("Token generated for employer: %s", emp.Email)

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

func (h *Handler) LoginEmployer(c *gin.Context) {
    var employerLoginReq EmployerLoginReq
    err := c.ShouldBindJSON(&employerLoginReq)
    if err != nil {
        log.Printf("Failed to parse employer login request: %v", err)
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid input data",
            Errors:  map[string]string{"request": "Unable to parse request body"},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    log.Printf("Employer login attempt: %s", employerLoginReq.Email)

    empRes, err := h.client.LoginEmployer(h.ctx, &pb.EmpLoginReq{
        Email:    employerLoginReq.Email,
        Password: employerLoginReq.Password,
    })
    if err != nil {
        log.Printf("Failed employer login for: %s, error: %v", employerLoginReq.Email, err)
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusUnauthorized,
            Message: "Failed to login",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    log.Printf("Employer login successful: %s", employerLoginReq.Email)

    res := toEmployer(empRes)
    token, err := token.GenerateToken(res.Email, res.ID, string(Employer))
    if err != nil {
        log.Printf("Failed to generate token for employer: %s, error: %v", res.Email, err)
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Failed to generate token",
            Errors:  map[string]string{"token": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    log.Printf("Token generated for employer: %s", res.Email)

    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Employer Profile retrieved successfully",
        Data: map[string]interface{}{
            "employer": res,
            "token":    token,
        },
    }
    c.JSON(succRes.Code, succRes)
}