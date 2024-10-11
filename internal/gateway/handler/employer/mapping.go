package employer

import pb "github.com/ahdaan67/jobportal/utils/pb/employer"

func toPBEmployerReq(e *CreateEmployerReq) *pb.CreateEmployerReq {
	return &pb.CreateEmployerReq{
		Name:     e.Name,
		Email:    e.Email,
		Password: e.Password,
		Phone:    e.Phone,
		Address:  e.Address,
		Country:  e.Country,
		Website:  e.Website,
	}
}

func toEmployer(e *pb.EmployerRes) *EmployerRes {
	return &EmployerRes{
		ID:      e.Id,
		Name:    e.Phone,
		Email:   e.Email,
		Country: e.Country,
		Address: e.Address,
		Phone:   e.Phone,
		Website: e.Website,
	}
}
