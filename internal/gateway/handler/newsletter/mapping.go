package newsletter

import (
	"strconv"

	pb "github.com/ahdaan67/jobportal/utils/pb/newsletter"
)

func toNewsLetterRes(n *pb.NewsLetterRes) *NewsLetterRes {
	return &NewsLetterRes{
		ID:         n.Id,
		EmployerID: n.EmployerId,
		Content:    n.Content,
		IsFree:     n.IsFree,
		Amount:     n.Amount,
	}
}

func ToSubscription(pbSub *pb.SubscriptionRes) *SubscriptionRes {
	if pbSub == nil {
		return nil
	}

	startDate := pbSub.Startdate.AsTime()
	endDate := pbSub.Enddate.AsTime()


	return &SubscriptionRes{
		ID:          pbSub.Id,
		JobSeekerID: pbSub.JobseekerId,
		NewLetterID: pbSub.NewletterId,
		StartDate:   startDate,
		EndDate:     endDate,
		Status:      pbSub.Status,
	}
}

func ToPayment(pbPay *pb.PaymentRes) *PaymentRes {
	if pbPay == nil {
		return nil
	}

	return &PaymentRes{
		ID:             pbPay.Id,
		SubscriptionID: pbPay.SubscriptionId,
		Amount:         pbPay.Amount,
		Status:         pbPay.Status,
		Date:           pbPay.Date.AsTime(),
	}
}

func ToRazorpay(pbRazor *pb.RazorpayRes) *RazorpayRes {
	if pbRazor == nil {
		return nil
	}

	return &RazorpayRes{
		ID:        pbRazor.Id,
		PaymentID: pbRazor.PaymentId,
		PayID:     pbRazor.PayId,
		OrderID:   pbRazor.OrderId,
		Signature: pbRazor.Signature,
	}
}

func ToSPR(pbSPR *pb.SPR) (*SubscriptionRes, *PaymentRes, *RazorpayRes) {
	if pbSPR == nil {
		return nil, nil, nil 
	}
	return ToSubscription(pbSPR.Subscirption), ToPayment(pbSPR.Payment), ToRazorpay(pbSPR.Razorpay)
}

func strtoInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}
