package newsletter

import (
	"context"
	"log"
	"net/http"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/response"
	pb "github.com/ahdaan67/jobportal/utils/pb/newsletter"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	ctx    context.Context
	client pb.NewsLetterClient
	cfg    config.Config
}

func NewHandler(client pb.NewsLetterClient, cfg config.Config) *Handler {
	return &Handler{
		ctx:    context.Background(),
		client: client,
		cfg:    cfg,
	}
}

func (h *Handler) GetNewsLetter(c *gin.Context) {
	nlid := c.Param("newsletterid")
	v, err := strtoInt64(nlid)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid newsletter ID",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}
	nl, err := h.client.GetNewsLetter(h.ctx, &pb.NLid{Id: v})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: "failed to get Newsletter",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}
	res := toNewsLetterRes(nl)
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "successfully got Newsletter",
		Data:    res,
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) ListNewsLetters(c *gin.Context) {
	nlr, err := h.client.ListNewsLetters(h.ctx, &emptypb.Empty{})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: "failed to list News Letters",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	var res []*NewsLetterRes
	for _, p := range nlr.GetNewletters() {
		res = append(res, toNewsLetterRes(p))
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "List of available Newsletter Service retrieved successfully",
		Data:    res,
	}
	c.JSON(succRes.Code, succRes)
}

func (h *Handler) AddSubscription(c *gin.Context) {
	var s SubscriptionReq
	err := c.ShouldBindJSON(&s)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"request": "Unable to parse request body"},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	spr, err := h.client.AddSubscription(h.ctx, &pb.SubscriptionReq{Jobseekerid: s.JobseekerID, Newsletterid: s.NewsLetterID})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to add subscription",
			Errors:  map[string]string{"subscription": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	if spr.Razorpay == nil && spr.Payment == nil && spr.Subscirption == nil {
		succRes := response.SuccessResponse{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Subscription reactivated successfully.",
			Data:    nil,
		}
		c.JSON(succRes.Code, succRes)
		return
	}

	sub, pay, razor := ToSPR(spr)
	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "List of available Newsletter Service retrieved successfully",
		Data: map[string]interface{}{
			"sub_id":            sub.ID,
			"newsletter_id":     sub.NewLetterID,
			"status":            sub.Status,
			"razorpay_order_id": razor.OrderID,
			"amount":            pay.Amount,
			"payment_status":    pay.Status,
			"payment_date":      pay.Date,
		},
	}

	c.JSON(succRes.Code, succRes)
}

func (h *Handler) RazorpayPayment(c *gin.Context) {
	su := c.Param("subscription")
	subid, err := strtoInt64(su)
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid subscription ID",
			Errors:  map[string]string{"subid": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	spr, err := h.client.GetSubscriptionAndPaymentDetails(h.ctx, &pb.Subid{Id: subid})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "failed to get subscription details.",
			Errors:  map[string]string{"subscription": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	sub, pay, razor := ToSPR(spr)
	var bt string
	if sub.Status == "active" {
		bt = "Renew With Razorpay"
		client := razorpay.NewClient(h.cfg.RazorpayKey, h.cfg.RazorpaySecret)
		razor.OrderID, err = CreateRazorpayOrder(client, float64(pay.Amount))
		if err != nil {
			errRes := response.ErrorResponse{
				Status:  "error",
				Code:    http.StatusBadRequest,
				Message: "Failed to create Razorpay order.",
				Errors:  map[string]string{"razorpay_order": err.Error()},
			}
			c.JSON(errRes.Code, errRes)
			return
		}

	} else if sub.Status == "inactive" {
		bt = "Pay With Razorpay"
	}

	c.HTML(http.StatusOK, "payment.html", gin.H{
		"sub_id":   sub.ID,
		"amount":   pay.Amount,
		"button":   bt,
		"razor_id": razor.OrderID,
	})
}

func (h *Handler) VerifyPayment(c *gin.Context) {
	log.Println("Starting payment verification process")

	subid, err := strtoInt64(c.Query("sub_id"))
	if err != nil {
		log.Printf("Error converting sub_id to int64: %v", err)
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid job ID",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	payid := c.Query("payid")
	orderid := c.Query("orderid")
	signature := c.Query("signature")
	status := c.Query("status")

	log.Printf("Received payment verification details - SubID: %d, PayID: %s, OrderID: %s, Signature: %s, Status: %s", subid, payid, orderid, signature, status)

	if status == "success" {
		log.Printf("Processing successful payment for SubID: %d", subid)
		_, err := h.client.UpdateSubscriptionAndPayment(h.ctx, &pb.UpdateSubscriptionAndPaymentReq{
			Subid: &pb.Subid{Id: subid},
			Razorpay: &pb.Razorpay{
				PayId:     payid,
				OrderId:   orderid,
				Signature: signature,
			},
		})
		if err != nil {
			log.Printf("Error updating subscription and payment: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal Server Error"})
			return
		}
		log.Printf("Successfully updated subscription and payment for SubID: %d", subid)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
	log.Println("Completed payment verification process")
}

func (h *Handler) CancelSubscription(c *gin.Context) {
	subid, err := strtoInt64(c.Param("subid"))
	if err != nil {
		log.Printf("Error converting sub_id to int64: %v", err)
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid job ID",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	_, err = h.client.CancelSubscription(h.ctx, &pb.Subid{Id: subid})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Failed to cancel subscription",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "successfully cancelled subcription",
		Data:    map[string]interface{}{},
	}

	c.JSON(succRes.Code, succRes)
}

func (h *Handler) GetSubscribers(c *gin.Context) {
	empid, err := strtoInt64(c.Param("employerid"))
	if err != nil {
		log.Printf("Error converting employer_id to int64: %v", err)
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid employer_id",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	nlid, err := strtoInt64(c.Param("newsletterid"))
	if err != nil {
		log.Printf("Error converting newsletterid to int64: %v", err)
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid newsletterid",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	subs, err := h.client.GetSubscribers(h.ctx, &pb.GetSubscribersReq{Empid: empid, Nlid: nlid})
	if err != nil {
		errRes := response.ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve subscribers",
			Errors:  map[string]string{"grpc": err.Error()},
		}
		c.JSON(errRes.Code, errRes)
		return
	}

	var sbs []*SubscriptionRes
	for _, v := range subs.Subs {
		sbs = append(sbs, ToSubscription(v))
	}

	succRes := response.SuccessResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Subscribers retrieved successfully",
		Data: map[string]interface{}{
			"subscribers": sbs,
			"active":      subs.Active,
			"inactive":    subs.Inactive,
			"canceled":    subs.Canceled,
			"expired":     subs.Expired,
		},
	}
	c.JSON(succRes.Code, succRes)
}
