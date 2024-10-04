package job

import (
	"strconv"

	pb "github.com/ahdaan67/jobportal/utils/pb/job"
)

func toPBJobReq(j *JobReq) *pb.JobReq {
	return &pb.JobReq{
		EmployerId:      j.EmployerID,
		Title:           j.Title,
		Employmenttype:  j.EmploymentType,
		Description:     j.Description,
		Location:        j.Location,
		Salary:          float32(j.Salary),
		Experiencelevel: j.ExperienceLevel,
	}
}

func toJobRes(j *pb.JobRes) *JobRes {
	return &JobRes{
		ID:              j.Id,
		EmployerID:      j.EmployerId,
		Title:           j.Title,
		EmploymentType:  j.Employmenttype,
		Description:     j.Description,
		Location:        j.Location,
		Salary:          float64(j.Salary),
		ExperienceLevel: j.Experiencelevel,
		PostedDate:      j.Postedat.AsTime(),
	}
}

func toApplyJob(j *pb.ApplyJobRes) *ApplyJobRes {
	return &ApplyJobRes{
		ID:          j.Id,
		JobseekerID: j.Jobseekerid,
		JobID:       j.Jobid,
		Resume:      j.Resume,
		Status:      j.Status,
		AppliedAt:   j.Appliedat.AsTime(),
	}
}

func strtoInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}
