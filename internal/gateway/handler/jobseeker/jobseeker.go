package jobseeker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	"github.com/ahdaan67/jobportal/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	ctx    context.Context
	client pb.JobSeekerClient
	cfg    config.Config
	logger  *logrus.Logger
	logFile *os.File
}

var sj CreateJobseekerReq
var otp string

type Role string

const (
	JobSeeker Role = "jobseeker"
)

func NewHandler(client pb.JobSeekerClient, cfg config.Config, logfile string) *Handler {
	logger, logFile := logging.InitLogrusLogger(logfile)
	return &Handler{
		ctx:    context.Background(),
		client: client,
		cfg:    cfg,
		logger:  logger,
		logFile: logFile,
	}
}


func (h *Handler) JobSeekerSignup(c *gin.Context) {
	h.logger.Info("Received request for Job Seeker Signup")

	err := c.ShouldBindJSON(&sj)
	if err != nil {
		h.logger.WithError(err).Error("Failed to bind JSON")
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
		h.logger.WithError(err).Error("Validation error during signup")
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
		h.logger.WithError(err).WithField("email", sj.Email).Error("Failed to send OTP")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to send OTP",
			Errors:  map[string]string{"request": "Please try again..."},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.WithField("email", sj.Email).Info("Successfully sent OTP")
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Successfully sent OTP",
		Data: map[string]interface{}{
			"to_email": sj.Email,
		},
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) VerifyOTP(c *gin.Context) {
	h.logger.Info("Received request to verify OTP")

	type OTPv struct {
		OtpCode string `json:"otp_code"`
	}

	var o OTPv
	err := c.ShouldBindJSON(&o)
	if err != nil {
		h.logger.WithError(err).Error("Failed to bind JSON for OTP verification")
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
		h.logger.WithField("email", sj.Email).Error("Invalid OTP provided")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusUnauthorized,
			Message: "Invalid OTP",
			Errors: map[string]string{
				"otp":     "The provided OTP is incorrect",
				"request": "Please enter the details again.",
			},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	if err := sj.Validate(); err != nil {
		h.logger.WithError(err).Error("Validation error during OTP verification")
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
		h.logger.WithError(err).WithField("jobseeker", sj).Error("Failed to create jobseeker")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to create jobseeker",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	res := toCreateJobseekerRes(jobseeker)

	token, err := token.GenerateToken(res.Email, res.ID, string(JobSeeker))
	if err != nil {
		h.logger.WithError(err).WithField("email", res.Email).Error("Failed to generate token")
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
			Errors:  map[string]string{"token": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	h.logger.WithField("email", res.Email).Info("Jobseeker signed up successfully")
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Jobseeker signed up successfully",
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
    logrus.Info("Received request to create Job Seeker Profile")

    var js JobSeekerProfileReq
    err := c.ShouldBindJSON(&js)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid input data",
            Errors:  map[string]string{"request": "Unable to parse request body"},
        }
        logrus.WithField("error", err.Error()).Error("Failed to bind JSON for Job Seeker Profile")
        c.JSON(errRes.Code, errRes)
        return
    }

    jobseeker, err := h.client.CreateJobSeekerProfile(h.ctx, toPBJobSeekerProfileReq(&js))
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Failed to create jobseeker profile",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        logrus.WithField("request", js).WithField("error", err.Error()).Error("Failed to create Job Seeker Profile")
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
    logrus.Infof("Jobseeker profile created successfully: %v", res.ID)
    c.JSON(succRes.Code, succRes)
}

func (h *Handler) GetJobseekers(c *gin.Context) {
    h.logger.Info("Received request to get Jobseekers")

    pbjs, err := h.client.GetJobseekers(h.ctx, &emptypb.Empty{})
    if err != nil {
        h.logger.WithError(err).Error("Failed to retrieve Jobseekers")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to get jobseekers",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    var jobseekers []*CreateJobseekerRes
    for _, j := range pbjs.Jobseekers {
        jobseekers = append(jobseekers, toCreateJobseekerRes(j))
    }

    h.logger.WithField("count", len(jobseekers)).Info("Successfully retrieved jobseekers")
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Jobseekers retrieved successfully",
        Data: map[string]interface{}{
            "jobseekers": jobseekers,
        },
    }
    c.JSON(succRes.Code, succRes)
}

func (h *Handler) GetJobseekerProfile(c *gin.Context) {
    h.logger.Info("Received request to get Jobseeker Profile")

    id := c.Param("id")
    v, err := strtoInt64(id)
    if err != nil {
        h.logger.WithError(err).WithField("id", id).Error("Invalid jobseeker ID provided")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid jobseeker ID",
            Errors:  map[string]string{"request": "Unable to parse jobseeker ID"},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    jpr, err := h.client.GetJobseekerProfile(h.ctx, &pb.JobSeekerID{Id: v})
    if err != nil {
        h.logger.WithError(err).WithField("jobseeker_id", v).Error("Failed to retrieve Jobseeker Profile")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to get jobseeker profile",
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

    h.logger.WithFields(logrus.Fields{
        "jobseeker_id": v,
        "profile_status": ps,
    }).Info("Jobseeker Profile retrieved successfully")
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Jobseeker Profile retrieved successfully",
        Data: map[string]interface{}{
            "jobseeker":      res,
            "profile_status": ps,
        },
    }
    c.JSON(succRes.Code, succRes)
}

func (h *Handler) LoginJobseeker(c *gin.Context) {
    h.logger.Info("Received request for Jobseeker Login")

    var js JSLoginReq
    err := c.ShouldBindJSON(&js)
    if err != nil {
        h.logger.WithError(err).Error("Failed to bind JSON for Jobseeker Login")
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
        h.logger.WithError(err).WithField("email", js.Email).Error("Failed to login jobseeker")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusUnauthorized,
            Message: "Failed to login",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    res := toCreateJobseekerRes(jcr)

    token, err := token.GenerateToken(res.Email, res.ID, string(JobSeeker))
    if err != nil {
        h.logger.WithError(err).WithField("email", res.Email).Error("Failed to generate token")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to generate token",
            Errors:  map[string]string{"token": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    h.logger.WithField("email", res.Email).Info("Jobseeker logged in successfully")
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Jobseeker logged in successfully",
        Data: map[string]interface{}{
            "jobseeker": res,
            "token":     token,
        },
    }
    c.JSON(succRes.Code, succRes)
}

func (h *Handler) FollowEmployers(c *gin.Context) {

    h.logger.Info("Received request to Follow Employer")

    id, exit := c.Get("id")
    if !exit {
        h.logger.Error("Failed to get id from authorization")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusUnauthorized,
            Message: "Failed to get id from authorization",
        }
        c.JSON(errRes.Code, errRes)
        return
    }
    jsid, ok := id.(int64)
    if !ok {
        h.logger.Error("Failed to convert id to int64")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to convert from any to int64",
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    var e FollowEmployerReq
    err := c.ShouldBindJSON(&e)
    e.JobseekerID = jsid
    if err != nil {
        h.logger.WithError(err).Error("Failed to bind JSON for Follow Employer")
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
        h.logger.WithError(err).WithFields(logrus.Fields{
            "jobseeker_id": e.JobseekerID,
            "employer_id": e.EmployerID,
        }).Error("Failed to follow employer")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to follow employer",
            Errors:  map[string]string{"employer": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    res := toPBEmployerRes(er)
    h.logger.WithFields(logrus.Fields{
        "jobseeker_id": e.JobseekerID,
        "employer_id": e.EmployerID,
    }).Info("Successfully followed employer")
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Successfully followed employer",
        Data: map[string]interface{}{
            "employer": res,
        },
    }
   c.JSON(succRes.Code, succRes)
}

func (h *Handler) UnFollowEmployers(c *gin.Context) {
    h.logger.Info("Received request to Unfollow Employer")

    id, exit := c.Get("id")
    if !exit {
        h.logger.Error("Failed to get id from authorization")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusUnauthorized,
            Message: "Failed to get id from authorization",
        }
        c.JSON(errRes.Code, errRes)
        return
    }
    jsid, ok := id.(int64)
    if !ok {
        h.logger.Error("Failed to convert id to int64")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to convert from any to int64",
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    var e FollowEmployerReq
    err := c.ShouldBindJSON(&e)
    e.JobseekerID = jsid
    if err != nil {
        h.logger.WithError(err).Error("Failed to bind JSON for Unfollow Employer")
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
        h.logger.WithError(err).WithFields(logrus.Fields{
            "jobseeker_id": e.JobseekerID,
            "employer_id": e.EmployerID,
        }).Error("Failed to unfollow employer")
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to unfollow employer",
            Errors:  map[string]string{"employer": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        return
    }

    h.logger.WithFields(logrus.Fields{
        "jobseeker_id": e.JobseekerID,
        "employer_id": e.EmployerID,
    }).Info("Successfully unfollowed employer")
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Successfully unfollowed employer",
        Data:    nil,
    }
    c.JSON(succRes.Code, succRes)
}

func (h *Handler) GetFollowingEmployers(c *gin.Context) {
	id, exit := c.Get("id")
	if !exit {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to get id from authorization",
		}
		c.JSON(errRes.Code, errRes)
		return
	}
	jsid, ok := id.(int64)
	if !ok {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to convert from any to int64",
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	pbemps, err := h.client.GetFollowingEmployers(h.ctx, &pb.JobSeekerID{Id: jsid})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to get following employers",
			Errors:  map[string]string{"following_employers": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	var res []*EmployerRes
	for _,e := range pbemps.Emp {
		res = append(res, toPBEmployerRes(e))
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Successfully retrieved following employers",
		Data:    res,
	}

	c.JSON(succRes.Code, succRes)
}