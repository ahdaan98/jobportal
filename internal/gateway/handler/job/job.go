package job

import (
	"context"
	"io"
	"net/http"
	"os"

	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	ResumeSavePath = "uploads/"

	Accept  = "accepted"
	Reject  = "rejected"
	Pending = "pending"
)

type Handler struct {
	ctx     context.Context
	client  pb.JobClient
	cfg     config.Config
	logger  *logrus.Logger
	logFile *os.File
}

func NewHandler(client pb.JobClient, cfg config.Config, logfile string) *Handler {
	logger, logFile := logging.InitLogrusLogger(logfile)
	return &Handler{
		ctx:    context.Background(),
		client: client,
		cfg:    cfg,
		logger:  logger,
		logFile: logFile,
	}
}

func (h *Handler) CreateJob(c *gin.Context) {
	h.logger.Info("CreateJob: Received request to create job")

	var j JobReq
	if err := c.ShouldBindJSON(&j); err != nil {
		h.logger.WithError(err).Error("CreateJob: Invalid input data")
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"request": "Unable to parse request body"},
		})
		return
	}

	job, err := h.client.CreateJob(h.ctx, toPBJobReq(&j))
	if err != nil {
		h.logger.WithError(err).Error("CreateJob: Failed to create job")
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to create job",
			Errors:  map[string]string{"grpc": err.Error()},
		})
		return
	}

	res := toJobRes(job)
	h.logger.Infof("CreateJob: Job created successfully with ID %d", res.ID)
	c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Job created successfully",
		Data:    res,
	})
}

func (h *Handler) GetJob(c *gin.Context) {
	h.logger.Info("GetJob: Received request to get job")

	id := c.Param("id")
	v, err := strtoInt64(id)
	if err != nil {
		h.logger.WithError(err).Error("GetJob: Invalid job ID")
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid job ID",
			Errors:  map[string]string{"grpc": err.Error()},
		})
		return
	}

	job, err := h.client.GetJob(h.ctx, &pb.GetJobReq{Id: v})
	if err != nil {
		h.logger.WithError(err).Error("GetJob: Job not found")
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: "Job not found",
			Errors:  map[string]string{"grpc": err.Error()},
		})
		return
	}

	res := toJobRes(job)
	h.logger.Infof("GetJob: Job retrieved successfully with ID %d", res.ID)
	c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Job retrieved successfully",
		Data:    res,
	})
}

func (h *Handler) ListJobs(c *gin.Context) {
	h.logger.Info("ListJobs: Received request to list jobs")

	ljr, err := h.client.ListJobs(h.ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.WithError(err).Error("ListJobs: Failed to list jobs")
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: "Failed to list Jobs",
			Errors:  map[string]string{"grpc": err.Error()},
		})
		return
	}

	var res []*JobRes
	for _, p := range ljr.GetJobs() {
		res = append(res, toJobRes(p))
	}

	h.logger.Infof("ListJobs: Successfully retrieved %d jobs", len(res))
	c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "List of Jobs retrieved successfully",
		Data:    res,
	})
}

func (h *Handler) UpdateJob(c *gin.Context) {
    id := c.Param("id")
    h.logger.Infof("Received request to update job with ID: %s", id)

    v, err := strtoInt64(id)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid job ID",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": id,
            "error":  errRes,
        }).Error("Invalid job ID error")
        return
    }

    var j JobReq
    err = c.ShouldBindJSON(&j)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid input data",
            Errors:  map[string]string{"request": "Unable to parse request body"},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": id,
            "error":  errRes,
        }).Error("Error binding JSON")
        return
    }

    h.logger.WithFields(logrus.Fields{
        "job_id": id,
        "job_request": j,
    }).Info("Received job request for update")

    job, err := h.client.UpdateJob(h.ctx, &pb.UpdateJobReq{Id: v, Job: toPBJobReq(&j)})
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to update job",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": id,
            "error":  errRes,
        }).Error("Error updating job")
        return
    }

    res := toJobRes(job)
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "Job updated successfully",
        Data:    res,
    }
    c.JSON(succRes.Code, succRes)
    h.logger.WithFields(logrus.Fields{
        "job_id": id,
        "response": succRes,
    }).Info("Job updated successfully")
}

func (h *Handler) ApplyJob(c *gin.Context) {
    h.logger.Info("Received request to apply for job")

    jsid := c.Param("jobseekerid")
    h.logger.Infof("Jobseeker ID from request: %s", jsid)

    jsid64, err := strtoInt64(jsid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid jobseeker ID",
            Errors:  map[string]string{"jobseeker_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "jobseeker_id": jsid,
            "error":        errRes,
        }).Error("Invalid jobseeker ID error")
        return
    }
    h.logger.Infof("Parsed jobseeker ID: %d", jsid64)

    jobid := c.Param("jobid")
    h.logger.Infof("Job ID from request: %s", jobid)

    jobid64, err := strtoInt64(jobid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid job ID",
            Errors:  map[string]string{"job_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid,
            "error":  errRes,
        }).Error("Invalid job ID error")
        return
    }
    h.logger.Infof("Parsed job ID: %d", jobid64)

    resume, err := c.FormFile("resume")
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Error getting resume file",
            Errors:  map[string]string{"form_file": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "error": errRes,
        }).Error("Error retrieving resume file")
        return
    }
    h.logger.Infof("Received resume file: %s", resume.Filename)

    file, err := resume.Open()
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to open resume file",
            Errors:  map[string]string{"resume": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "resume_file": resume.Filename,
            "error":       errRes,
        }).Error("Error opening resume file")
        return
    }
    defer file.Close()
    h.logger.Info("Opened resume file successfully")

    fileBytes, err := io.ReadAll(file)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to read resume file",
            Errors:  map[string]string{"resume": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "resume_file": resume.Filename,
            "error":       errRes,
        }).Error("Error reading resume file")
        return
    }
    h.logger.Info("Read resume file successfully")

    Path := ResumeSavePath + resume.Filename
    err = c.SaveUploadedFile(resume, Path)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Error saving resume",
            Errors:  map[string]string{"resume": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "resume_file": resume.Filename,
            "error":       errRes,
        }).Error("Error saving resume file")
        return
    }
    h.logger.Infof("Saved resume file to path: %s", Path)

    apj, err := h.client.ApplyJob(h.ctx, &pb.ApplyJobReq{Jobid: jobid64, Jobseekerid: jsid64, Resumedata: fileBytes})
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusInternalServerError,
            Message: "Failed to apply for job",
            Errors:  map[string]string{"grpc": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id":       jobid,
            "jobseeker_id": jsid,
            "error":        errRes,
        }).Error("Error applying for job")
        return
    }
    h.logger.Info("Successfully applied for job")

    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusCreated,
        Message: "Applied for job successfully",
        Data:    apj,
    }
    c.JSON(succRes.Code, succRes)
    h.logger.WithFields(logrus.Fields{
        "job_id":       jobid,
        "jobseeker_id": jsid,
        "response":     succRes,
    }).Info("Applied for job successfully and responded")
}


func (h *Handler) ListApplyJobsByid(c *gin.Context) {
    h.logger.Info("Received request to list applicants by job ID")

    jobid := c.Param("jobid")
    h.logger.Infof("Job ID from request: %s", jobid)

    jobid64, err := strtoInt64(jobid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid job ID",
            Errors:  map[string]string{"job_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid,
            "error":  errRes,
        }).Error("Invalid job ID error")
        return
    }
    h.logger.Infof("Parsed job ID: %d", jobid64)

    laj, err := h.client.ListApplyJobByID(h.ctx, &pb.GetJobReq{Id: jobid64})
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Failed to get applicants",
            Errors:  map[string]string{"applicants": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid64,
            "error":  errRes,
        }).Error("Error retrieving applicants")
        return
    }

    var res []*ApplyJobRes
    for _, p := range laj.GetApplyjobs() {
        res = append(res, toApplyJob(p))
    }

    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "List of applicants retrieved successfully",
        Data:    res,
    }
    c.JSON(succRes.Code, succRes)
    h.logger.Infof("Successfully retrieved applicants for job ID %d", jobid64)
}

func (h *Handler) AcceptApplicant(c *gin.Context) {
    h.logger.Info("Received request to accept applicant")

    jsid := c.Param("jobseekerid")
    h.logger.Infof("Jobseeker ID from request: %s", jsid)

    jsid64, err := strtoInt64(jsid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid jobseeker ID",
            Errors:  map[string]string{"jobseeker_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "jobseeker_id": jsid,
            "error":        errRes,
        }).Error("Invalid jobseeker ID error")
        return
    }
    h.logger.Infof("Parsed jobseeker ID: %d", jsid64)

    jobid := c.Param("jobid")
    h.logger.Infof("Job ID from request: %s", jobid)

    jobid64, err := strtoInt64(jobid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid job ID",
            Errors:  map[string]string{"job_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid,
            "error":  errRes,
        }).Error("Invalid job ID error")
        return
    }
    h.logger.Infof("Parsed job ID: %d", jobid64)

    ua := UpdateApplicants{
        JobID:       jobid64,
        JobseekerID: jsid64,
        Status:      Accept,
    }

    aj, err := h.client.UpdateApplicants(h.ctx, &pb.UpdateAppReq{Jobid: ua.JobID, Jobseekerid: ua.JobseekerID, Status: ua.Status})
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "failure",
            Code:    http.StatusBadRequest,
            Message: "Failed to accept applicant",
            Errors: map[string]string{
                "update_error": err.Error(),
            },
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id":       jobid64,
            "jobseeker_id": jsid64,
            "error":        errRes,
        }).Error("Error accepting applicant")
        return
    }

    res := toApplyJob(aj)
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusAccepted,
        Message: "Applicant status updated successfully",
        Data: map[string]interface{}{
            "applicant": res,
        },
    }

    c.JSON(http.StatusAccepted, succRes)
    h.logger.Infof("Successfully accepted applicant with ID %d for job ID %d", jsid64, jobid64)
}

func (h *Handler) RejectApplicant(c *gin.Context) {
    h.logger.Info("Received request to reject applicant")

    jsid := c.Param("jobseekerid")
    h.logger.Infof("Jobseeker ID from request: %s", jsid)

    jsid64, err := strtoInt64(jsid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid jobseeker ID",
            Errors:  map[string]string{"jobseeker_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "jobseeker_id": jsid,
            "error":        errRes,
        }).Error("Invalid jobseeker ID error")
        return
    }
    h.logger.Infof("Parsed jobseeker ID: %d", jsid64)

    jobid := c.Param("jobid")
    h.logger.Infof("Job ID from request: %s", jobid)

    jobid64, err := strtoInt64(jobid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid job ID",
            Errors:  map[string]string{"job_id": err.Error()},
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid,
            "error":  errRes,
        }).Error("Invalid job ID error")
        return
    }
    h.logger.Infof("Parsed job ID: %d", jobid64)

    ua := UpdateApplicants{
        JobID:       jobid64,
        JobseekerID: jsid64,
        Status:      Reject,
    }

    aj, err := h.client.UpdateApplicants(h.ctx, &pb.UpdateAppReq{Jobid: ua.JobID, Jobseekerid: ua.JobseekerID, Status: ua.Status})
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "failure",
            Code:    http.StatusBadRequest,
            Message: "Failed to reject applicant",
            Errors: map[string]string{
                "update_error": err.Error(),
            },
        }
        c.JSON(errRes.Code, errRes)
        h.logger.WithFields(logrus.Fields{
            "job_id":       jobid64,
            "jobseeker_id": jsid64,
            "error":        errRes,
        }).Error("Error rejecting applicant")
        return
    }

    res := toApplyJob(aj)
    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusAccepted,
        Message: "Applicant status updated successfully",
        Data: map[string]interface{}{
            "applicant": res,
        },
    }

    c.JSON(http.StatusAccepted, succRes)
    h.logger.Infof("Successfully rejected applicant with ID %d for job ID %d", jsid64, jobid64)
}

func (h *Handler) GetApplicantsByStatus(c *gin.Context) {
    jobid := c.Param("jobid")
    h.logger.Infof("Received request to get applicants for job ID: %s", jobid)

    jobid64, err := strtoInt64(jobid)
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "Invalid jobs ID",
            Errors:  map[string]string{"job_id": err.Error()},
        }
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid,
            "error":  err.Error(),
        }).Error("Invalid job ID error")
        c.JSON(errRes.Code, errRes)
        return
    }

    status := c.Param("status")
    validStatus := map[string]bool{
        Accept:  true,
        Reject:  true,
        Pending: true,
    }

    if _, exists := validStatus[status]; !exists {
        errRes := response.ErrorResponse{
            Status:  "failure",
            Code:    http.StatusBadRequest,
            Message: "Invalid status provided. Accepted values are: Accept, Reject, Pending.",
            Errors:  map[string]string{"status": status},
        }
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid64,
            "status": status,
            "error":  "Invalid status",
        }).Error("Invalid status error")
        c.JSON(errRes.Code, errRes)
        return
    }

    ljr, err := h.client.Getapplicantsbystatus(h.ctx, &pb.GetApplicantReq{Jobid: jobid64, Status: status})
    if err != nil {
        errRes := response.ErrorResponse{
            Status:  "error",
            Code:    http.StatusBadRequest,
            Message: "failed to get applicants",
            Errors:  map[string]string{"applicants": err.Error()},
        }
        h.logger.WithFields(logrus.Fields{
            "job_id": jobid64,
            "status": status,
            "error":  err.Error(),
        }).Error("Error retrieving applicants by status")
        c.JSON(errRes.Code, errRes)
        return
    }

    var res []*ApplyJobRes
    for _, p := range ljr.GetApplyjobs() {
        res = append(res, toApplyJob(p))
    }

    succRes := response.SuccessResponse{
        Status:  "success",
        Code:    http.StatusOK,
        Message: "List of " + status + " applicants retrieved successfully",
        Data:    res,
    }
    h.logger.Infof("Successfully retrieved applicants by status '%s' for job ID %d", status, jobid64)
    c.JSON(succRes.Code, succRes)
}
