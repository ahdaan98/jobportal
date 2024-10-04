package jobseeker

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	"github.com/ahdaan67/jobportal/utils/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Handler struct {
	ctx    context.Context
	client pb.JobSeekerClient
	cfg    config.Config
}

var sj CreateJobseekerReq
var otp string

type Role string

const (
	JobSeeker Role = "jobseeker"
)

func NewHandler(client pb.JobSeekerClient, cfg config.Config) *Handler {
	return &Handler{
		ctx:    context.Background(),
		client: client,
		cfg:    cfg,
	}
}

func (h *Handler) JobSeekerSignup(c *gin.Context) {
	err := c.ShouldBindJSON(&sj)
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

	if err := sj.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Errors:  map[string]string{"validation": err.Error()},
		})
		return
	}

	otp, err = SendOTP(sj.Email, h.cfg)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to sent otp",
			Errors:  map[string]string{"request": "Please try again..."},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "successfully sented otp",
		Data: map[string]interface{}{
			"to_email": sj.Email,
		},
	}

	c.JSON(succRes.Code, succRes)
}

func (h *Handler) VerifyOTP(c *gin.Context) {
	type OTPv struct {
		OtpCode string `json:"otp_code"`
	}

	var o OTPv
	err := c.ShouldBindJSON(&o)
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

	if !VerifyOTP(otp, sj.Email) {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid otp",
			Errors: map[string]string{
				"otp":     "The provided otp is incorrect",
				"request": "Please enter the details again.",
			},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	if err := sj.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Errors:  map[string]string{"request": "Please enter the details again."},
		})
		return
	}

	jobseeker, err := h.client.CreateJobseeker(h.ctx, toPBCreateJobseeker(&sj))
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to create jobseeker",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toCreateJobseekerRes(jobseeker)

	token, err := token.GenerateToken(res.Email, res.ID, string(JobSeeker))
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
		Message: "Jobseeker Signed Up successfully",
		Data: map[string]interface{}{
			"jobseeker": res,
			"token":     token,
			"next_step": "setup profile",
		},
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) LinkedinSignIn(c *gin.Context) {
	Linkedinconfig := func(cfg config.Config) oauth2.Config {
		LinkedinConfig := oauth2.Config{
			RedirectURL:  "https://oauth.pstmn.io/v1/callback",
			ClientID:     cfg.LinkedinClientID,
			ClientSecret: cfg.LinkedinClientSecretID,
			Scopes:       []string{"openid", "profile", "email"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
				TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
			},
		}

		return LinkedinConfig
	}(h.cfg)

	url := Linkedinconfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusSeeOther, url)
}

func (h *Handler) LinkedInSignInCallback(c *gin.Context) {
	accessToken := c.Query("access_token")
	if accessToken == "" {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"access_token": "missing access token"},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	resp, err := func(token string) (*http.Response, error) {
		url := "https://api.linkedin.com/v2/userinfo"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token)
		client := &http.Client{}
		return client.Do(req)
	}(accessToken)

	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Failed to retrieve user info from LinkedIn API",
			Errors:  map[string]string{"api_error": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Unexpected response from LinkedIn API, status code: %d", resp.StatusCode),
			Errors:  map[string]string{"linkedin_api": "Failed to fetch user data from LinkedIn API"},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	js, err := io.ReadAll(resp.Body)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to read response body from LinkedIn API",
			Errors:  map[string]string{"response_body": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}
	c.JSON(http.StatusOK, string(js))
}

func (h *Handler) CreateJobSeekerProfile(c *gin.Context) {
	var js JobSeekerProfileReq
	err := c.ShouldBindJSON(&js)
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

	jobseeker, err := h.client.CreateJobSeekerProfile(h.ctx, toPBJobSeekerProfileReq(&js))
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to create jobseeker profile",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toProfileJobRes(jobseeker)
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Jobseeker SetUp profile successfully",
		Data: map[string]interface{}{
			"profile": res,
		},
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) GetJobseekerProfile(c *gin.Context) {
	id := c.Param("id")
	v, err := strtoInt64(id)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobseeker ID",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	jpr, err := h.client.GetJobseekerProfile(h.ctx, &pb.JobSeekerID{Id: v})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to get jobseeker profile",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toProfileJobRes(jpr)

	var ps string
	if res.Summary == "" || res.Education == "" || res.City == "" || res.Country == "" {
		ps = "incomplete"
	} else {
		ps = "completed"
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Jobseeker Profile retrieved successfully",
		Data: map[string]interface{}{
			"jobseeker":      res,
			"profile_status": ps,
		},
	}

	c.JSON(succRes.Code, succRes)
}

func (h *Handler) LoginJobseeker(c *gin.Context) {
	var js JSLoginReq
	err := c.ShouldBindJSON(&js)
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

	jcr, err := h.client.LoginJobseeker(h.ctx, &pb.JSLoginReq{Email: js.Email, Password: js.Password})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to login",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toCreateJobseekerRes(jcr)

	token, err := token.GenerateToken(res.Email, res.ID, string(JobSeeker))
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
		Message: "Jobseeker Profile retrieved successfully",
		Data: map[string]interface{}{
			"jobseeker": res,
			"token":     token,
		},
	}

	c.JSON(succRes.Code, succRes)
}

func (h *Handler) FollowEmployers(c *gin.Context) {
	var e FollowEmployerReq
	err := c.ShouldBindJSON(&e)
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

	er, err := h.client.FollowEmployer(h.ctx, &pb.FollowEmployerReq{Jobseekerid: e.JobseekerID, Employerid: e.EmployerID})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to follow employer",
			Errors:  map[string]string{"employer": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toPBEmployerRes(er)
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Successfully followed employer",
		Data: map[string]interface{}{
			"employer": res,
		},
	}

	c.JSON(succRes.Code, succRes)
}

func (h *Handler) UnFollowEmployers(c *gin.Context) {
	var e FollowEmployerReq
	err := c.ShouldBindJSON(&e)
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

	_, err = h.client.UnFollowEmployer(h.ctx, &pb.FollowEmployerReq{Jobseekerid: e.JobseekerID, Employerid: e.EmployerID})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to follow employer",
			Errors:  map[string]string{"employer": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Successfully unfollowed employer",
		Data:    nil,
	}

	c.JSON(succRes.Code, succRes)
}
