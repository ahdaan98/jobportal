package newsletter

import "time"

type NewsLetterRes struct {
	ID         int64   `json:"id"`
	EmployerID int64   `json:"employer_id"`
	Content    string  `json:"content"`
	IsFree     bool    `json:"is_free"`
	Amount     float32 `json:"amount"`
}

type SubscriptionReq struct {
	JobseekerID  int64 `json:"jobseeker_id"`
	NewsLetterID int64 `json:"newletter_id"`
}

type SubscriptionRes struct {
	ID          int64     `json:"id"`
	JobSeekerID int64     `json:"jobseeker_id"`
	NewLetterID int64     `json:"newletter_id"`
	StartDate   time.Time `json:"startdate"`
	EndDate     time.Time `json:"enddate"`
	Status      string    `json:"status"`
}

type PaymentRes struct {
	ID             int64     `json:"id"`
	SubscriptionID int64     `json:"subscription_id"`
	Amount         float32   `json:"amount"`
	Status         string    `json:"status"`
	Date           time.Time `json:"date"`
}

type RazorpayRes struct {
	ID        int64  `json:"id"`
	PaymentID int64  `json:"payment_id"`
	PayID     string `json:"pay_id"`
	OrderID   string `json:"order_id"`
	Signature string `json:"signature"`
}
