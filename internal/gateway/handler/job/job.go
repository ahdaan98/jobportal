package job

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	ResumeSavePath = "uploads/"

	Accept  = "accepted"
	Reject  = "rejected"
	Pending = "pending"
)

type Handler struct {
	ctx    context.Context
	client pb.JobClient
	cfg    config.Config
}

func NewHandler(client pb.JobClient, cfg config.Config) *Handler {
	return &Handler{
		ctx:    context.Background(),
		client: client,
		cfg:    cfg,
	}
}

func (h *Handler) CreateJob(c *gin.Context) {
	var j JobReq
	err := c.ShouldBindJSON(&j)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"request": "Unable to parse request body"},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("CreateJob: %v", errRes)
		return
	}

	job, err := h.client.CreateJob(h.ctx, toPBJobReq(&j))
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to create job",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("CreateJob: %v", errRes)
		return
	}

	res := toJobRes(job)
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Job created successfully",
		Data:    res,
	}
	c.JSON(succRes.Code, succRes)
	log.Printf("CreateJob: %v", succRes)
}

func (h *Handler) GetJob(c *gin.Context) {
	id := c.Param("id")
	v, err := strtoInt64(id)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid job ID",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("GetJob: %v", errRes)
		return
	}

	job, err := h.client.GetJob(h.ctx, &pb.GetJobReq{Id: v})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: "Job not found",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("GetJob: %v", errRes)
		return
	}

	res := toJobRes(job)
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Job retrieved successfully",
		Data:    res,
	}
	c.JSON(succRes.Code, succRes)
	log.Printf("GetJob: %v", succRes)
}

func (h *Handler) ListJobs(c *gin.Context) {
	ljr, err := h.client.ListJobs(h.ctx, &emptypb.Empty{})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: "failed to list Jobs",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ListJobs: %v", errRes)
		return
	}

	var res []*JobRes
	for _, p := range ljr.GetJobs() {
		res = append(res, toJobRes(p))
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "List of Jobs retrieved successfully",
		Data:    res,
	}
	c.JSON(succRes.Code, succRes)
	log.Printf("ListJobs: %v", succRes)
}

func (h *Handler) UpdateJob(c *gin.Context) {
	id := c.Param("id")
	v, err := strtoInt64(id)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid job ID",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("UpdateJob: %v", errRes)
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
		log.Printf("UpdateJob: %v", errRes)
		return
	}

	log.Println("UpdateJob: Received job request:", j)

	job, err := h.client.UpdateJob(h.ctx, &pb.UpdateJobReq{Id: v, Job: toPBJobReq(&j)})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to update job",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("UpdateJob: %v", errRes)
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
	log.Printf("UpdateJob: %v", succRes)
}

func (h *Handler) ApplyJob(c *gin.Context) {
	jsid := c.Param("jobseekerid")
	jsid64, err := strtoInt64(jsid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobseeker ID",
			Errors:  map[string]string{"jobseeker_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ApplyJob: %v", errRes)
		return
	}

	jobid := c.Param("jobid")
	jobid64, err := strtoInt64(jobid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobs ID",
			Errors:  map[string]string{"job_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ApplyJob: %v", errRes)
		return
	}

	resume, err := c.FormFile("resume")
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "error getting resume file",
			Errors:  map[string]string{"form_file": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ApplyJob: %v", errRes)
		return
	}

	file, err := resume.Open()
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to open resume file",
			Errors:  map[string]string{"resume": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ApplyJob: %v", errRes)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to read resume file",
			Errors:  map[string]string{"resume": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ApplyJob: %v", errRes)
		return
	}

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
		log.Printf("ApplyJob: %v", errRes)
		return
	}

	apj, err := h.client.ApplyJob(h.ctx, &pb.ApplyJobReq{Jobid: jobid64, Jobseekerid: jsid64, Resumedata: fileBytes})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to apply for job",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		log.Printf("ApplyJob: %v", errRes)
		return
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Applied for job successfully",
		Data:    apj,
	}
	c.JSON(succRes.Code, succRes)
	log.Printf("ApplyJob: %v", succRes)
}

func (h *Handler) ListApplyJobsByid(c *gin.Context) {
	jobid := c.Param("jobid")
	jobid64, err := strtoInt64(jobid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid job ID",
			Errors:  map[string]string{"job_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	laj, err := h.client.ListApplyJobByID(h.ctx, &pb.GetJobReq{Id: jobid64})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to get applicants",
			Errors:  map[string]string{"applicants": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
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
}

func (h *Handler) AcceptApplicant(c *gin.Context) {
	jsid := c.Param("jobseekerid")
	jsid64, err := strtoInt64(jsid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobseeker ID",
			Errors:  map[string]string{"jobseeker_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	jobid := c.Param("jobid")
	jobid64, err := strtoInt64(jobid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobs ID",
			Errors:  map[string]string{"job_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

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
			Message: "failed to accept applicant",
			Errors: map[string]string{
				"update_error": err.Error(),
			},
		}
		c.JSON(errRes.Code, errRes)
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
}

func (h *Handler) RejectApplicant(c *gin.Context) {
	jsid := c.Param("jobseekerid")
	jsid64, err := strtoInt64(jsid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobseeker ID",
			Errors:  map[string]string{"jobseeker_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	jobid := c.Param("jobid")
	jobid64, err := strtoInt64(jobid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobs ID",
			Errors:  map[string]string{"job_id": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

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
			Message: "failed to accept applicant",
			Errors: map[string]string{
				"update_error": err.Error(),
			},
		}
		c.JSON(errRes.Code, errRes)
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
}

func (h *Handler) GetApplicantsByStatus(c *gin.Context) {
	jobid := c.Param("jobid")
	jobid64, err := strtoInt64(jobid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid jobs ID",
			Errors:  map[string]string{"job_id": err.Error()},
		}
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
		Message: "List of "+status+" applicants retrieved successfully",
		Data:    res,
	}
	c.JSON(succRes.Code, succRes)
}
