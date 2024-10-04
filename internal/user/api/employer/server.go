package employer

import (
	"context"

	"github.com/ahdaan67/jobportal/internal/user/storer/employer"
	pb "github.com/ahdaan67/jobportal/utils/pb/employer"
)

type Server struct {
	storer *employer.EMPLOYERstorer
	pb.UnimplementedEmployerServer
}

func NewServer(storer *employer.EMPLOYERstorer) *Server {
	return &Server{
		storer: storer,
	}
}

func (s *Server) CreateEmployer(ctx context.Context, req *pb.CreateEmployerReq) (*pb.EmployerRes, error) {
	id, err := s.storer.CreateEmployer(ctx, toCreateEmployer(req))
	if err != nil {
		return nil, err
	}

	emp, err := s.storer.GetEmployer(ctx, id)
	if err != nil {
		return nil, err
	}

	return toEmployerRes(emp), nil
}