package jobseeker

import (
	"github.com/ahdaan67/jobportal/internal/user/storer/employer"
	"github.com/ahdaan67/jobportal/internal/user/storer/jobseeker"
	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toCreateJobseeker(j *pb.CreateJobseekerReq) *jobseeker.CreateJobseekerReq {
	return &jobseeker.CreateJobseekerReq{
		FirstName:   j.Firstname,
		LastName:    j.Lastname,
		Email:       j.Email,
		Password:    j.Password,
		Gender:      j.Gender,
		Phone:       j.Phone,
		DateOfBirth: j.Dateofbirth.AsTime(),
	}
}

func toPBCreateJobseeker(j *jobseeker.CreateJobseekerRes) *pb.CreateJobseekerRes {
	return &pb.CreateJobseekerRes{
		Id:          j.ID,
		Firstname:   j.FirstName,
		Lastname:    j.LastName,
		Email:       j.Email,
		Gender:      j.Gender,
		Phone:       j.Phone,
		Dateofbirth: timestamppb.New(j.DateOfBirth),
		Createdat:   timestamppb.New(j.CreatedAt),
	}
}

func toProfileJobseeker(j *pb.JobSeekerProfileReq) *jobseeker.JobSeekerProfileReq {
	return &jobseeker.JobSeekerProfileReq{
		JobseekerID: j.Jobseekerid,
		Summary:     j.Summary,
		City:        j.City,
		Country:     j.Country,
		Education:   j.Education,
		Experience:  j.Experience,
	}
}

func toPBJobseekerProfile(j *jobseeker.JobSeekerProfileRes) *pb.JobSeekerProfileRes {
	js := toPBCreateJobseeker(&j.CreateJobseekerRes)
	return &pb.JobSeekerProfileRes{
		Jobseeker:  js,
		Summary:    j.Summary,
		City:       j.City,
		Country:    j.Country,
		Education:  j.Education,
		Experience: j.Experience,
	}
}

func toEmployerRes(e *employer.EmployerRes) *pb.EmployerRes {
	return &pb.EmployerRes{
		Id:      e.ID,
		Name:    e.Phone,
		Country: e.Country,
		Address: e.Address,
		Phone:   e.Phone,
		Website: e.Website,
	}
}
