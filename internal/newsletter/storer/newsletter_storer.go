package storer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	Active    = "active"
	InActive  = "inactive"
	Success   = "success"
	Failure   = "failure"
	Pending   = "pending"
	Cancelled = "canceled"
	Expired   = "expired"
)

type NEWSLETTERstorer struct {
	db *sqlx.DB
}

func NewNEWSLETTERStorer(db *sqlx.DB) *NEWSLETTERstorer {
	return &NEWSLETTERstorer{
		db: db,
	}
}

func (ns *NEWSLETTERstorer) ListnewsLetters(ctx context.Context) ([]*NewsLetterRes, error) {
	var n []*NewsLetterRes
	err := ns.db.SelectContext(ctx, &n, "SELECT * FROM employers_newsletter")
	if err != nil {
		return nil, fmt.Errorf("error listing newsletters: %w", err)
	}

	return n, nil
}

func (ns *NEWSLETTERstorer) GetNewsLetter(ctx context.Context, newsletterid int64) (*NewsLetterRes, error) {
	var n NewsLetterRes
	err := ns.db.GetContext(ctx, &n, "SELECT * FROM employers_newsletter WHERE id=$1", newsletterid)
	if err != nil {
		return nil, fmt.Errorf("error getting newsletter: %w", err)
	}

	return &n, nil
}

func (ns *NEWSLETTERstorer) CreateSubscription(ctx context.Context, s *SubscriptionReq) (int64, error) {
	stmt, err := ns.db.PrepareNamedContext(ctx, "INSERT INTO subscriptions(jobseeker_id, newletter_id) VALUES (:jobseeker_id, :newletter_id) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("error preparing subscription insert statement: %w", err)
	}

	var id int64
	err = stmt.Get(&id, s)
	if err != nil {
		return 0, fmt.Errorf("error executing subscription insert: %w", err)
	}
	return id, nil
}

func (ns *NEWSLETTERstorer) IsSubscriptionActive(ctx context.Context, s *SubscriptionReq) (bool, error) {
	var count int64
	err := ns.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM subscriptions WHERE jobseeker_id=$1 AND newletter_id=$2 AND status=$3", s.JobseekerID, s.NewsLetterID, Active)
	if err != nil {
		return false, fmt.Errorf("error getting subscription active: %w", err)
	}

	return count > 0, nil
}

func (ns *NEWSLETTERstorer) IsSubscriptionInActive(ctx context.Context, s *SubscriptionReq) (bool, error) {
	var count int64
	err := ns.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM subscriptions WHERE jobseeker_id=$1 AND newletter_id=$2 AND status=$3", s.JobseekerID, s.NewsLetterID, InActive)
	if err != nil {
		return false, fmt.Errorf("error getting subscription active: %w", err)
	}

	return count > 0, nil
}

func (ns *NEWSLETTERstorer) CreatePayment(ctx context.Context, p *PaymentReq) (int64, error) {
	stmt, err := ns.db.PrepareNamedContext(ctx, "INSERT INTO payments(subscription_id, amount) VALUES (:subscription_id, :amount) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("error preparing payment insert statement: %w", err)
	}

	var id int64
	err = stmt.Get(&id, p)
	if err != nil {
		return 0, fmt.Errorf("error executing payment insert: %w", err)
	}

	return id, nil
}

func (ns *NEWSLETTERstorer) CreateRazorpayOrder(ctx context.Context, r *RazorpayReq) (int64, error) {
	stmt, err := ns.db.PrepareNamedContext(ctx, "INSERT INTO razorpay_details(payment_id, order_id) VALUES (:payment_id, :order_id) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("error preparing payment insert statement: %w", err)
	}

	var id int64
	err = stmt.Get(&id, r)
	if err != nil {
		return 0, fmt.Errorf("error executing payment insert: %w", err)
	}

	return id, nil
}

func (ns *NEWSLETTERstorer) GetSubscription(ctx context.Context, id int64) (*SubscriptionRes, error) {
	var s SubscriptionRes
	err := ns.db.GetContext(ctx, &s, "SELECT * FROM subscriptions WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting subscription by id: %w", err)
	}

	return &s, nil
}

func (ns *NEWSLETTERstorer) GetSubscriptionbyJobseekerandNewsletterid(ctx context.Context, s *SubscriptionReq) (*SubscriptionRes, error) {
	var sr SubscriptionRes
	err := ns.db.GetContext(ctx, &sr, "SELECT * FROM subscriptions WHERE jobseeker_id=$1 AND newletter_id=$2", s.JobseekerID, s.NewsLetterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting subscription by id: %w", err)
	}

	return &sr, nil
}

func (ns *NEWSLETTERstorer) GetPayment(ctx context.Context, id int64) (*PaymentRes, error) {
	var p PaymentRes
	err := ns.db.GetContext(ctx, &p, "SELECT * FROM payments WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting payment by id: %w", err)
	}

	return &p, nil
}

func (ns *NEWSLETTERstorer) GetPaymentBysubid(ctx context.Context, id int64) (*PaymentRes, error) {
	var p PaymentRes
	err := ns.db.GetContext(ctx, &p, "SELECT * FROM payments WHERE subscription_id = $1 ORDER BY id DESC LIMIT 1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting payment by sub id: %w", err)
	}

	return &p, nil
}

func (ns *NEWSLETTERstorer) GetRazorpay(ctx context.Context, id int64) (*RazorpayRes, error) {
	var r RazorpayRes
	err := ns.db.GetContext(ctx, &r, "SELECT * FROM razorpay_details WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting razorpay details by id: %w", err)
	}

	return &r, nil
}

func (ns *NEWSLETTERstorer) GetRazorpayBypaymentid(ctx context.Context, id int64) (*RazorpayRes, error) {
	var r RazorpayRes
	err := ns.db.GetContext(ctx, &r, "SELECT * FROM razorpay_details WHERE payment_id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting razorpay details by id: %w", err)
	}

	return &r, nil
}

func (ns *NEWSLETTERstorer) UpdateRazorpayPayidAndSignature(ctx context.Context, r *Razorpay) error {
	_, err := ns.db.NamedExecContext(ctx, "UPDATE razorpay_details SET pay_id=:pay_id, signature=:signature  WHERE order_id=:order_id", r)
	if err != nil {
		return fmt.Errorf("error updating razorpay pay_id and signature by id: %w", err)
	}

	return nil
}

func (ns *NEWSLETTERstorer) UpdatePaymentStatus(ctx context.Context, paymentID int64, status string) error {
	_, err := ns.db.ExecContext(ctx, "UPDATE payments SET status=$1 WHERE id=$2", status, paymentID)
	if err != nil {
		return fmt.Errorf("error updating payment status: %w", err)
	}
	return nil
}

func (ns *NEWSLETTERstorer) UpdateSubscriptionStatus(ctx context.Context, subscriptionID int64, status string) error {
	_, err := ns.db.ExecContext(ctx, "UPDATE subscriptions SET status=$1 WHERE id=$2", status, subscriptionID)
	if err != nil {
		return fmt.Errorf("error updating subscription status: %w", err)
	}
	return nil
}

func (ns *NEWSLETTERstorer) IsStartAndEnddateExist(ctx context.Context, id int64) (bool, error) {
	query := `SELECT COUNT(*) FROM subscriptions WHERE id=$1 AND startdate IS NOT NULL AND enddate IS NOT NULL;`
	var count int
	err := ns.db.GetContext(ctx, &count, query, id)
	if err != nil {
		return false, fmt.Errorf("error checking existence: %w", err)
	}
	return count > 0, nil
}

func (ns *NEWSLETTERstorer) AddSubStartAndEnddate(ctx context.Context, id int64) error {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 1, 0)
	log.Printf("Adding subscription: id=%d, startDate=%v, endDate=%v", id, startDate, endDate)
	query := `
        UPDATE subscriptions 
        SET startdate = $1, enddate = $2 
        WHERE id = $3;
        `

	_, err := ns.db.ExecContext(ctx, query, startDate, endDate, id)
	if err != nil {
		return fmt.Errorf("error updating subscription start and end date: %w", err)
	}

	return nil
}

func (ns *NEWSLETTERstorer) UpdateEnddate(ctx context.Context, id int64) error {
	query := `
	SELECT startdate, enddate FROM subscriptions WHERE id=$1;
	`
	var s SubscriptionDate
	err := ns.db.GetContext(ctx, &s, query, id)
	if err != nil {
		return fmt.Errorf("error getting subscription start and end date: %w", err)
	}

	updateQuery := `
	UPDATE subscriptions 
	SET enddate = enddate + INTERVAL '1 month' 
	WHERE id = $1;
	`
	_, err = ns.db.ExecContext(ctx, updateQuery, id)
	if err != nil {
		return fmt.Errorf("error updating enddate: %w", err)
	}

	return nil
}

func (ns *NEWSLETTERstorer) IsSubscriptionExpired(ctx context.Context, id int64) (bool, error) {
	query := `
	SELECT enddate FROM subscriptions WHERE id=$1;
	`
	var endDate time.Time
	err := ns.db.GetContext(ctx, &endDate, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error getting subscription end date: %w", err)
	}

	if endDate.After(time.Now()) {
		return true, nil
	}

	return false, nil
}

func (ns *NEWSLETTERstorer) IsEmployerNewsLetterExist(ctx context.Context, empid, nlid int64) (bool, error) {
	var count int64
	err := ns.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM employers_newsletter WHERE employer_id=$1 AND id=$2", empid, nlid)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error getting employer newsletter : %w", err)
	}

	return count > 0, nil
}

func (ns *NEWSLETTERstorer) GetNewsLetterSubscribers(ctx context.Context, id int64) ([]*SubscriptionRes, error) {
	var subs []*SubscriptionRes
	err := ns.db.SelectContext(ctx, &subs, "SELECT * FROM subscriptions WHERE newletter_id=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no subscribers")
		}
		return nil, fmt.Errorf("failed to get subscribers")
	}

	return subs, nil
}

func (ns *NEWSLETTERstorer) GetCountOfSubscribersByStatus(ctx context.Context, id int64, status string) (int64, error) {
	var count int64
	err := ns.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM subscriptions WHERE newletter_id=$1 AND status=$2", id, status)
	if err != nil {
		return 0, fmt.Errorf("error getting count of %v subscribers: %w", status, err)
	}

	return count, nil
}
