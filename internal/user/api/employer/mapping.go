package employer

import (
	"github.com/ahdaan67/jobportal/internal/user/storer/employer"
	pb "github.com/ahdaan67/jobportal/utils/pb/employer"
)

func toCreateEmployer(e *pb.CreateEmployerReq) *employer.CreateEmployerReq {
	return &employer.CreateEmployerReq{
		Name:    e.Name,
		Email:   e.Email,
		Phone:   e.Phone,
		Address: e.Address,
		Country: e.Country,
		Website: e.Website,
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