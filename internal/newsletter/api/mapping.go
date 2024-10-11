package api

import (
	"database/sql"

	"github.com/ahdaan67/jobportal/internal/newsletter/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/newsletter"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPBNewsLetterRes(n *storer.NewsLetterRes) *pb.NewsLetterRes {
	return &pb.NewsLetterRes{
		Id:         n.ID,
		EmployerId: n.EmployerID,
		Content:    n.Content,
		IsFree:     n.IsFree,
		Amount:     n.Amount,
	}
}

func toPBSubscription(s *storer.SubscriptionRes) *pb.SubscriptionRes {
	return &pb.SubscriptionRes{
		Id:          s.ID,
		JobseekerId: s.JobSeekerID,
		NewletterId: s.NewLetterID,
		Startdate:   nullTimeToPB(s.StartDate),
		Enddate:     nullTimeToPB(s.EndDate),
		Status:      s.Status,
	}
}

func toPBPayment(p *storer.PaymentRes) *pb.PaymentRes {
	return &pb.PaymentRes{
		Id:             p.ID,
		SubscriptionId: p.SubscriptionID,
		Amount:         p.Amount,
		Status:         p.Status,
		Date:           timestamppb.New(p.Date),
	}
}

func toPBRazorpay(r *storer.RazorpayRes) *pb.RazorpayRes {
	return &pb.RazorpayRes{
		Id:        r.ID,
		PaymentId: r.PaymentID,
		PayId:     r.PayID.String,
		OrderId:   r.OrderID.String,
		Signature: r.Signature.String,
	}
}

func toPBSPR(s *storer.SubscriptionRes, p *storer.PaymentRes, r *storer.RazorpayRes) *pb.SPR {
	return &pb.SPR{
		Subscirption: toPBSubscription(s),
		Payment:      toPBPayment(p),
		Razorpay:     toPBRazorpay(r),
	}
}

func nullTimeToPB(nt sql.NullTime) *timestamppb.Timestamp {
	if nt.Valid {
		return timestamppb.New(nt.Time)
	}
	return nil
}
