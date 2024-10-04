package jobseeker

import (
	"context"
	"fmt"
	"log"

	"github.com/ahdaan67/jobportal/internal/user/storer/employer"
	"github.com/ahdaan67/jobportal/internal/user/storer/jobseeker"
	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	storer  *jobseeker.JOBSEEKERstorer
	estorer *employer.EMPLOYERstorer
	pb.UnimplementedJobSeekerServer
}

func NewServer(storer *jobseeker.JOBSEEKERstorer, estorer *employer.EMPLOYERstorer) *Server {
	return &Server{
		storer: storer,
		estorer: estorer,
	}
}

func (s *Server) CreateJobseeker(ctx context.Context, req *pb.CreateJobseekerReq) (*pb.CreateJobseekerRes, error) {
	jr, err := s.storer.CreateJobseeker(ctx, toCreateJobseeker(req))
	if err != nil {
		return nil, err
	}

	return toPBCreateJobseeker(jr), nil
}

func (s *Server) CreateJobSeekerProfile(ctx context.Context, req *pb.JobSeekerProfileReq) (*pb.JobSeekerProfileRes, error) {
	err := s.storer.CreateJobseekerProfile(ctx, toProfileJobseeker(req))
	if err != nil {
		return nil, err
	}

	jpr, err := s.storer.GetJobSeeker(ctx, req.GetJobseekerid())
	if err != nil {
		return nil, err
	}

	return toPBJobseekerProfile(jpr), nil
}

func (s *Server) GetJobseekerProfile(ctx context.Context, req *pb.JobSeekerID) (*pb.JobSeekerProfileRes, error) {
	jpr, err := s.storer.GetJobSeeker(ctx, req.GetId())
	if err != nil {
		if err.Error() == NoRowExist {
			js, err := s.storer.GetBasicJSProfilebyID(ctx, req.GetId())
			if err != nil {
				return nil, fmt.Errorf("error getting job seeker profile: %w", err)
			}

			return &pb.JobSeekerProfileRes{Jobseeker: toPBCreateJobseeker(js)}, nil
		} else {
			log.Printf("Error retrieving jobseeker: %v", err)
			return nil, err
		}
	}

	return toPBJobseekerProfile(jpr), nil
}

func (s *Server) LoginJobseeker(ctx context.Context, req *pb.JSLoginReq) (*pb.CreateJobseekerRes, error) {
	ok, err := s.storer.IsJSexist(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("no jobseeker exist")
	}

	d, err := s.storer.GetJSpass(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if req.Password != d.Pass {
		return nil, fmt.Errorf("invalid Email or Password")
	}

	jp, err := s.storer.GetBasicJSProfile(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if jp.Password != req.Password {
		return nil, fmt.Errorf("invalid Email or Password")
	}

	return toPBCreateJobseeker(jp), nil
}

func (s *Server) FollowEmployer(ctx context.Context, req *pb.FollowEmployerReq) (*pb.EmployerRes, error) {
	err := s.storer.FollowEmployer(ctx, &jobseeker.FollowEmployerReq{JobseekerID: req.Jobseekerid, EmployerID: req.Employerid})
	if err != nil {
		return nil, err
	}

	er, err := s.estorer.GetEmployer(ctx, req.Employerid)
	if err != nil {
		return nil, err
	}

	res := toEmployerRes(er)
	return res, nil
}

func (s *Server) UnFollowEmployer(ctx context.Context, req *pb.FollowEmployerReq) (*emptypb.Empty, error) {
	err := s.storer.UnFollowEmployer(ctx, &jobseeker.FollowEmployerReq{JobseekerID: req.Jobseekerid, EmployerID: req.Employerid})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}