package employer

import (
	"context"
	"fmt"

	"github.com/ahdaan67/jobportal/internal/user/storer/employer"
	pb "github.com/ahdaan67/jobportal/utils/pb/employer"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Server) LoginEmployer(ctx context.Context, req *pb.EmpLoginReq) (*pb.EmployerRes, error) {
    ok, err := s.storer.IsEmployerExist(ctx, req.Email)
    if err != nil {
        return nil, err
    }

    if !ok {
        return nil, fmt.Errorf("no employer exists")
    }

    // Fetch employer's stored password
    d, err := s.storer.GetEmployerPass(ctx, req.Email)
    if err != nil {
        return nil, err
    }

    if req.Password != d.Pass {
        return nil, fmt.Errorf("invalid Email or Password")
    }

    ep, err := s.storer.GetEmployer(ctx, d.Id)
    if err != nil {
        return nil, err
    }

    return toEmployerRes(ep), nil
}

func (s *Server) GetEmployers(ctx context.Context, req *emptypb.Empty) (*pb.EmployerDetails, error) {
    var employers pb.EmployerDetails

    emps, err := s.storer.GetEmployers(ctx)
    if err != nil {
        return nil, err
    }

    for _, e := range emps {
        emmp := &pb.EmployerDetail{
            Employer:       toEmployerRes(e),
            Followers:      0,
            Jobopertunities: 0,
        }

        emmp.Followers, err = s.storer.GetFollowersCount(ctx, e.ID)
        if err != nil {
            return nil, err
        }

        emmp.Jobopertunities, err = s.storer.GetJobCounts(ctx, e.ID)
        if err != nil {
            return nil, err
        }

        employers.Employer = append(employers.Employer, emmp)
    }

    return &employers, nil
}