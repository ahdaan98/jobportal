package api

import (
	"context"
	"fmt"
	"log"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/newsletter/service"
	"github.com/ahdaan67/jobportal/internal/newsletter/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/newsletter"
	"github.com/razorpay/razorpay-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	storer *storer.NEWSLETTERstorer
	cfg    config.Config
	pb.UnimplementedNewsLetterServer
}

func NewServer(storer *storer.NEWSLETTERstorer, cfg config.Config) *Server {
	return &Server{
		storer: storer,
		cfg:    cfg,
	}
}

func (s *Server) CreateNewsletter(ctx context.Context, req *pb.NewsLetterReq) (*pb.NewsLetterRes, error) {
	newsletter := toNewsletter(req)
	if newsletter.Content == "" {
		return nil, fmt.Errorf("please provide an newsletter content description")
	}

	if newsletter.Amount <= 0 {
		return nil, fmt.Errorf("cannot use -ve amount")
	}

	if newsletter.IsFree && newsletter.Amount != 0{
		return nil, fmt.Errorf("cannot enter amount for free service")
	}
	
	if !newsletter.IsFree&& newsletter.Amount == 0{
		return nil, fmt.Errorf("select free if no amount needed")
	}

	ns, err := s.storer.CreateNewsletter(ctx, newsletter)
	if err!= nil {
		return nil, err
	}

	return toPBNewsLetterRes(ns), nil
}

func (s *Server) GetSubscription(ctx context.Context, req *pb.SubscriptionReq) (*pb.ArrSPR, error) {
	subids, err := s.storer.GetSubscriptionIds(ctx, req.Jobseekerid)
	if err != nil {
		return nil, err
	}

	var sprs []*pb.SPR
	for _,subid := range subids {

		sub, err := s.storer.GetSubscription(ctx, subid)
		if err != nil {
			log.Printf("Error retrieving subscription: %v", err)
			return nil, err
		}

		pay, err := s.storer.GetPaymentBysubid(ctx, sub.ID)
		if err != nil {
			log.Printf("Error retrieving payment: %v", err)
			return nil, err
		}

		razor, err := s.storer.GetRazorpayBypaymentid(ctx, pay.ID)
		if err != nil {
			log.Printf("Error retrieving Razorpay details: %v", err)
			return nil, err
		}

		sprs=append(sprs,toPBSPR(sub, pay, razor))
	}
	log.Println(sprs)
	return &pb.ArrSPR{Spr: sprs}, nil
}

func (s *Server) GetNewsLetter(ctx context.Context, req *pb.NLid) (*pb.NewsLetterRes, error) {
	nl, err := s.storer.GetNewsLetter(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return toPBNewsLetterRes(nl), err
}

func (s *Server) ListNewsLetters(ctx context.Context, req *emptypb.Empty) (*pb.ListNewsLetterRes, error) {
	lns, err := s.storer.ListnewsLetters(ctx)
	if err != nil {
		return nil, err
	}

	lnr := make([]*pb.NewsLetterRes, 0, len(lns))
	for _, ln := range lns {
		lnr = append(lnr, toPBNewsLetterRes(ln))
	}

	return &pb.ListNewsLetterRes{Newletters: lnr}, nil
}

func (s *Server) AddSubscription(ctx context.Context, req *pb.SubscriptionReq) (*pb.SPR, error) {
	log.Println("Starting AddSubscription process")
	
	// Retrieve the subscription for the given jobseeker and newsletter
	subb, err := s.storer.GetSubscriptionbyJobseekerandNewsletterid(ctx, &storer.SubscriptionReq{
		JobseekerID: req.Jobseekerid,
		NewsLetterID: req.Newsletterid,
	})
	if err != nil {
		log.Printf("Error retrieving subscription: %v", err)
		return nil, err
	}

	// If no subscription exists, create a new one
	if subb == nil {
		log.Printf("No existing subscription found for JobseekerID: %d and NewsletterID: %d. Creating new subscription.", req.Jobseekerid, req.Newsletterid)

		log.Printf("Retrieving newsletter information for NewsletterID: %d", req.Newsletterid)
		nl, err := s.storer.GetNewsLetter(ctx, req.Newsletterid)
		if err != nil {
			log.Printf("Error retrieving newsletter: %v", err)
			return nil, err
		}

		if nl == nil {
			return nil, fmt.Errorf("no newsletter found for NewsletterID %d", req.Newsletterid)
		}

		log.Printf("Creating new subscription for JobseekerID: %d, NewsletterID: %d", req.Jobseekerid, req.Newsletterid)
		subid, err := s.storer.CreateSubscription(ctx, &storer.SubscriptionReq{
			JobseekerID: req.Jobseekerid,
			NewsLetterID: req.Newsletterid,
		})
		if err != nil {
			log.Printf("Error creating subscription: %v", err)
			return nil, err
		}

		log.Printf("Creating payment for SubscriptionID: %d", subid)
		paymentid, err := s.storer.CreatePayment(ctx, &storer.PaymentReq{
			SubscriptionID: subid,
			Amount:        nl.Amount,
		})
		if err != nil {
			log.Printf("Error creating payment: %v", err)
			return nil, err
		}

		log.Println("Creating Razorpay order")
		client := razorpay.NewClient(s.cfg.RazorpayKey, s.cfg.RazorpaySecret)
		razororderid, err := service.CreateRazorpayOrder(client, float64(nl.Amount))
		if err != nil {
			log.Printf("Error creating Razorpay order: %v", err)
			return nil, err
		}

		log.Printf("Storing Razorpay order with OrderID: %s", razororderid)
		razorid, err := s.storer.CreateRazorpayOrder(ctx, &storer.RazorpayReq{
			PaymentID: paymentid,
			OrderID:   razororderid,
		})
		if err != nil {
			log.Printf("Error storing Razorpay order: %v", err)
			return nil, err
		}

		log.Println("Fetching subscription, payment, and Razorpay details")
		sub, err := s.storer.GetSubscription(ctx, subid)
		if err != nil {
			log.Printf("Error retrieving subscription: %v", err)
			return nil, err
		}

		pay, err := s.storer.GetPayment(ctx, paymentid)
		if err != nil {
			log.Printf("Error retrieving payment: %v", err)
			return nil, err
		}

		razor, err := s.storer.GetRazorpay(ctx, razorid)
		if err != nil {
			log.Printf("Error retrieving Razorpay details: %v", err)
			return nil, err
		}

		log.Println("Successfully completed AddSubscription process")
		return toPBSPR(sub, pay, razor), nil
	}

	// Handle existing subscriptions
	log.Printf("Existing subscription found: %v", subb)
	switch subb.Status {
	case storer.Cancelled:
		yes, err := s.storer.IsSubscriptionExpired(ctx, subb.ID)
		if err != nil {
			log.Printf("Error checking if subscription is expired: %v", err)
			return nil, err
		}
		if yes {
			err = s.storer.UpdateSubscriptionStatus(ctx, subb.ID, storer.Active)
			if err != nil {
				log.Printf("Error activating subscription: %v", err)
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("subscription is already initiated - sub_status=%v - sub_id=%v, please renew the payment", subb.Status, subb.ID)
		}
		return &pb.SPR{Subscirption: nil, Razorpay: nil, Payment: nil}, nil

	case storer.Active, storer.Expired:
		return nil, fmt.Errorf("subscription is already initiated - sub_status=%v - sub_id=%v, please renew the payment", subb.Status, subb.ID)

	case storer.InActive:
		return nil, fmt.Errorf("subscription is already initiated - subscription is %v - sub_id=%v, due to payment pending", subb.Status, subb.ID)
	}

	return nil, fmt.Errorf("unhandled subscription status: %v", subb.Status)
}

func (s *Server) GetSubscriptionAndPaymentDetails(ctx context.Context, req *pb.Subid) (*pb.SPR, error) {
	log.Println("Starting GetSubscriptionAndPaymentDetails process")

	sub, err := s.storer.GetSubscription(ctx, req.Id)
	if err != nil {
		log.Printf("Error retrieving subscription with ID: %d, %v", req.Id, err)
		return nil, err
	}

	log.Printf("Fetching payment details for SubscriptionID: %d", sub.ID)
	pay, err := s.storer.GetPaymentBysubid(ctx, sub.ID)
	if err != nil {
		log.Printf("Error retrieving payment by subscription ID: %v", err)
		return nil, err
	}

	log.Printf("Fetching Razorpay details for PaymentID: %d", pay.ID)
	razor, err := s.storer.GetRazorpayBypaymentid(ctx, pay.ID)
	if err != nil {
		log.Printf("Error retrieving Razorpay details: %v", err)
		return nil, err
	}

	log.Println("Successfully retrieved subscription, payment, and Razorpay details")
	return toPBSPR(sub, pay, razor), nil
}

func (s *Server) UpdateSubscriptionAndPayment(ctx context.Context, req *pb.UpdateSubscriptionAndPaymentReq) (*emptypb.Empty, error) {
	log.Println("Starting UpdateSubscriptionAndPayment process")

	p, err := s.storer.GetPaymentBysubid(ctx, req.Subid.Id)
	if err != nil {
		log.Printf("Error retrieving payment by subscription ID: %v", err)
		return &emptypb.Empty{}, err
	}

	log.Printf("Payment status: %s", p.Status)

	if p.Status != storer.Pending {
		sub, err := s.storer.GetSubscription(ctx, req.Subid.Id)
		if err != nil {
			log.Printf("Error retrieving subscription: %v", err)
			return nil, err
		}

		log.Printf("Creating new payment for subscription ID: %d", sub.ID)
		paymentid, err := s.storer.CreatePayment(ctx, &storer.PaymentReq{SubscriptionID: sub.ID, Amount: p.Amount})
		if err != nil {
			log.Printf("Error creating new payment: %v", err)
			return nil, err
		}

		log.Println("Initiating Razorpay client and creating order")
		client := razorpay.NewClient(s.cfg.RazorpayKey, s.cfg.RazorpaySecret)
		razororderid, err := service.CreateRazorpayOrder(client, float64(p.Amount))
		if err != nil {
			log.Printf("Error creating Razorpay order: %v", err)
			return nil, err
		}

		log.Printf("Razorpay order created with ID: %s", razororderid)
		_, err = s.storer.CreateRazorpayOrder(ctx, &storer.RazorpayReq{PaymentID: paymentid, OrderID: razororderid})
		if err != nil {
			log.Printf("Error storing Razorpay order: %v", err)
			return nil, err
		}

		log.Println("Updating Razorpay payment ID and signature")
		err = s.storer.UpdateRazorpayPayidAndSignature(ctx, &storer.Razorpay{
			PayID:     req.Razorpay.PayId,
			OrderID:   razororderid,
			Signature: req.Razorpay.Signature,
		})
		if err != nil {
			log.Printf("Error updating Razorpay payment ID and signature: %v", err)
			return nil, err
		}

		log.Println("Updating payment status to Success")
		err = s.storer.UpdatePaymentStatus(ctx, paymentid, storer.Success)
		if err != nil {
			log.Printf("Error updating payment status: %v", err)
			return nil, err
		}
	}

	if p.Status == storer.Pending {
		log.Println("Updating Razorpay payment ID and signature for pending payment")
		err = s.storer.UpdateRazorpayPayidAndSignature(ctx, &storer.Razorpay{
			PayID:     req.Razorpay.PayId,
			OrderID:   req.Razorpay.OrderId,
			Signature: req.Razorpay.Signature,
		})
		if err != nil {
			log.Printf("Error updating Razorpay payment ID and signature for pending payment: %v", err)
			return nil, err
		}

		log.Println("Updating payment status to Success")
		err = s.storer.UpdatePaymentStatus(ctx, p.ID, storer.Success)
		if err != nil {
			log.Printf("Error updating payment status: %v", err)
			return nil, err
		}
	}

	log.Println("Checking if start and end date exist for subscription")
	exist, err := s.storer.IsStartAndEnddateExist(ctx, req.Subid.Id)
	if err != nil {
		log.Printf("Error checking start and end date existence: %v", err)
		return nil, err
	}

	if exist {
		log.Println("Updating end date for subscription")
		err = s.storer.UpdateEnddate(ctx, req.Subid.Id)
	} else {
		log.Println("Adding start and end date for subscription")
		err = s.storer.AddSubStartAndEnddate(ctx, req.Subid.Id)
	}

	if err != nil {
		log.Printf("Error updating or adding subscription dates: %v", err)
		return nil, err
	}

	log.Println("Updating subscription status to Active")
	err = s.storer.UpdateSubscriptionStatus(ctx, req.Subid.Id, storer.Active)
	if err != nil {
		log.Printf("Error updating subscription status: %v", err)
		return nil, err
	}

	log.Println("Successfully completed UpdateSubscriptionAndPayment process")
	return &emptypb.Empty{}, nil
}

func (s *Server) CancelSubscription(ctx context.Context, req *pb.Subid) (*emptypb.Empty, error) {
	sub, err := s.storer.GetSubscription(ctx, req.Id)
	if err != nil {
		log.Printf("Error retrieving subscription: %v", err)
		return nil, err
	}

	if sub.Status == storer.Cancelled {
		return nil, fmt.Errorf("subscription already cancelled")
	}

	if sub.Status == storer.InActive {
		return nil, fmt.Errorf("subscription is inactive and cannot be cancelled")
	}

	if sub.Status == storer.Active {
		err = s.storer.UpdateSubscriptionStatus(ctx, req.Id, storer.Cancelled)
		if err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetSubscribers(ctx context.Context, req *pb.GetSubscribersReq) (*pb.GetSubscribersRes, error) {
	exist, err := s.storer.IsEmployerNewsLetterExist(ctx, req.Empid, req.Nlid)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("no newsletter existz")
	}

	subs, err := s.storer.GetNewsLetterSubscribers(ctx, req.Nlid)
	if err != nil {
		return nil, err
	}

	ljr := make([]*pb.SubscriptionRes, 0, len(subs))
	for _, lj := range subs {
		ljr = append(ljr, toPBSubscription(lj))
	}

	active, err := s.storer.GetCountOfSubscribersByStatus(ctx, req.Nlid, storer.Active)
	if err != nil {
		return nil, err
	}

	inactive, err := s.storer.GetCountOfSubscribersByStatus(ctx, req.Nlid, storer.InActive)
	if err != nil {
		return nil, err
	}

	canceled, err := s.storer.GetCountOfSubscribersByStatus(ctx, req.Nlid, storer.Cancelled)
	if err != nil {
		return nil, err
	}

	expired, err := s.storer.GetCountOfSubscribersByStatus(ctx, req.Nlid, storer.Expired)
	if err != nil {
		return nil, err
	}

	return &pb.GetSubscribersRes{Subs: ljr, Active: active, Inactive: inactive, Canceled: canceled, Expired: expired}, nil
}
