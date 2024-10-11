package jobseeker

import (
	"strconv"

	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPBCreateJobseeker(j *CreateJobseekerReq) *pb.CreateJobseekerReq {
	return &pb.CreateJobseekerReq{
		Firstname:   j.FirstName,
		Lastname:    j.LastName,
		Email:       j.Email,
		Password:    j.Password,
		Gender:      j.Gender,
		Phone:       j.Phone,
		Dateofbirth: timestamppb.New(j.DateOfBirth),
	}
}

func toCreateJobseekerRes(j *pb.CreateJobseekerRes) *CreateJobseekerRes {
	return &CreateJobseekerRes{
		ID:          j.Id,
		FirstName:   j.Firstname,
		LastName:    j.Lastname,
		Email:       j.Email,
		Gender:      j.Gender,
		Phone:       j.Phone,
		DateOfBirth: j.Dateofbirth.AsTime(),
		CreatedAt:   j.Createdat.AsTime(),
	}
}

func toPBJobSeekerProfileReq(j *JobSeekerProfileReq) *pb.JobSeekerProfileReq {
	return &pb.JobSeekerProfileReq{
		Jobseekerid: j.JobseekerID,
		Summary:     j.Summary,
		City:        j.City,
		Country:     j.Country,
		Education:   j.Education,
		Experience:  j.Experience,
	}
}

func toProfileJobRes(j *pb.JobSeekerProfileRes) *JobSeekerProfileRes {
	return &JobSeekerProfileRes{
		CreateJobseekerRes: *toCreateJobseekerRes(j.Jobseeker),
		Summary:            j.Summary,
		City:               j.City,
		Country:            j.Country,
		Education:          j.Education,
		Experience:         j.Experience,
	}
}

func toPBEmployerRes(e *pb.EmployerRes) *EmployerRes {
	return &EmployerRes{
		ID:      e.Id,
		Name:    e.Phone,
		Country: e.Country,
		Address: e.Address,
		Phone:   e.Phone,
		Website: e.Website,
	}
}

func strtoInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err!=nil {
		return 0, err
	}

	return v, nil
}