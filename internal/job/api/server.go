package api

import (
	"context"
	"fmt"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/job/service"
	"github.com/ahdaan67/jobportal/internal/job/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	storer *storer.JOBstorer
	cfg    config.Config
	pb.UnimplementedJobServer
}

func NewServer(storer *storer.JOBstorer, cfg config.Config) *Server {
	return &Server{
		storer: storer,
		cfg:    cfg,
	}
}

func (s *Server) CreateJob(ctx context.Context, req *pb.JobReq) (*pb.JobRes, error) {
	if err := CheckZero(int64(req.Salary)); err != nil {
		return nil, err
	}

	jr, err := s.storer.CreateJob(ctx, toStorerJob(req))
	if err != nil {
		return nil, err
	}

	return toPBJobRes(jr), nil
}

func (s *Server) GetJob(ctx context.Context, req *pb.GetJobReq) (*pb.JobRes, error) {
	if err := CheckZero(req.Id); err != nil {
		return nil, err
	}

	jr, err := s.storer.GetJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return toPBJobRes(jr), nil
}

func (s *Server) ListJobs(ctx context.Context, req *emptypb.Empty) (*pb.ListJobRes, error) {
	ljs, err := s.storer.ListJobs(ctx)
	if err != nil {
		return nil, err
	}

	ljr := make([]*pb.JobRes, 0, len(ljs))
	for _, lj := range ljs {
		ljr = append(ljr, toPBJobRes(lj))
	}

	return &pb.ListJobRes{Jobs: ljr}, nil
}

func (s *Server) UpdateJob(ctx context.Context, req *pb.UpdateJobReq) (*pb.JobRes, error) {
	if err := CheckZero(req.Id); err != nil {
		return nil, err
	}

	if err := CheckZero(int64(req.Job.Salary)); err != nil {
		return nil, err
	}

	job, err := s.storer.GetJob(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	patchJobReq(job, req.Job)
	jr, err := s.storer.UpdateJob(ctx, req.Id, toStorerJob(req.Job))
	if err != nil {
		return nil, err
	}

	return toPBJobRes(jr), nil
}

func (s *Server) ApplyJob(ctx context.Context, req *pb.ApplyJobReq) (*pb.ApplyJobRes, error) {
	ok, err := s.storer.IsApplyJobexist(ctx, req.Jobseekerid, req.Jobid)
	if err != nil {
		return nil, err
	}

	if ok {
		return nil, fmt.Errorf("already applied")
	}

	fileUID := uuid.New()
	fileName := fileUID.String()

	awsimg, err := service.AddImageToAwsS3(s.cfg, req.Resumedata, fileName)
	if err != nil {
		return nil, err
	}

	req.Resume = awsimg

	err = s.storer.AddApplyJob(ctx, toApplyJob(req))
	if err != nil {
		return nil, err
	}

	aj, err := s.storer.GetApplyJob(ctx, toApplyJob(req))
	if err != nil {
		return nil, err
	}

	return toPBApplyJobRes(aj), nil
}

func (s *Server) ListApplyJobByID(ctx context.Context, req *pb.GetJobReq) (*pb.ListApplyJob, error) {
	ljs, err := s.storer.ListApplyJobbyID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	ljr := make([]*pb.ApplyJobRes, 0, len(ljs))
	for _, lj := range ljs {
		ljr = append(ljr, toPBApplyJobRes(lj))
	}

	return &pb.ListApplyJob{Applyjobs: ljr}, nil
}

func (s *Server) UpdateApplicants(ctx context.Context, req *pb.UpdateAppReq) (*pb.ApplyJobRes, error) {
	err := s.storer.UpdateApplicants(ctx, &storer.UpdateApplicants{JobID: req.Jobid, JobseekerID: req.Jobseekerid, Status: req.Status})
	if err != nil {
		return nil, err
	}

	aj, err := s.storer.GetApplyJob(ctx, toApplyJob(&pb.ApplyJobReq{Jobid: req.Jobid, Jobseekerid: req.Jobseekerid}))
	if err != nil {
		return nil, err
	}

	return toPBApplyJobRes(aj), nil
}

func (s *Server) Getapplicantsbystatus(ctx context.Context, req *pb.GetApplicantReq) (*pb.ListApplyJob, error) {
	ljs, err := s.storer.GetApplicantsByStatus(ctx, &storer.ApplicantsReq{JobID: req.Jobid, Status: req.Status})
	if err != nil {
		return nil, err
	}

	ljr := make([]*pb.ApplyJobRes, 0, len(ljs))
	for _, lj := range ljs {
		ljr = append(ljr, toPBApplyJobRes(lj))
	}

	return &pb.ListApplyJob{Applyjobs: ljr}, nil
}
