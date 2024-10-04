package api

import (
	"fmt"

	"github.com/ahdaan67/jobportal/internal/job/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toStorerJob(j *pb.JobReq) *storer.JobReq {
	return &storer.JobReq{
		EmployerID:      j.EmployerId,
		Title:           j.Title,
		EmploymentType:  j.Employmenttype,
		Description:     j.Description,
		Location:        j.Location,
		Salary:          float64(j.Salary),
		ExperienceLevel: j.Experiencelevel,
	}
}

func toPBJobRes(j *storer.JobRes) *pb.JobRes {
	return &pb.JobRes{
		Id:              j.ID,
		EmployerId:      j.EmployerID,
		Title:           j.Title,
		Employmenttype:  j.EmploymentType,
		Description:     j.Description,
		Location:        j.Location,
		Salary:          float32(j.Salary),
		Experiencelevel: j.ExperienceLevel,
		Postedat:        timestamppb.New(j.PostedDate),
	}
}

func patchJobReq(job *storer.JobRes, p *pb.JobReq) {
	if p.Title == "" {
		p.Title = job.Title
	}

	if p.Description == "" {
		p.Description = job.Description
	}

	if p.Employmenttype == "" {
		p.Employmenttype = job.EmploymentType
	}

	if p.Experiencelevel == "" {
		p.Experiencelevel = job.ExperienceLevel
	}

	if p.Location == "" {
		p.Location = job.Location
	}

	if p.Salary <= 0 {
		p.Salary = float32(job.Salary)
	}
}

func toApplyJob(j *pb.ApplyJobReq) *storer.ApplyJobReq {
	return &storer.ApplyJobReq{
		JobID:       j.Jobid,
		JobseekerID: j.Jobseekerid,
		Resume:      j.Resume,
	}
}

func toPBApplyJobRes(j *storer.ApplyJobRes) *pb.ApplyJobRes {
	return &pb.ApplyJobRes{
		Id:          j.ID,
		Jobid:       j.JobID,
		Jobseekerid: j.JobseekerID,
		Resume:      j.Resume,
		Status:      j.Status,
		Appliedat:   timestamppb.New(j.AppliedAt),
	}
}

func CheckZero(v int64) error {
	if v <= 0 {
		return fmt.Errorf("values cannot be zero or negative")
	}
	return nil
}
