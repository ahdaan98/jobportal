package storer

import (
	"database/sql"
	"time"
)

type NewsLetterRes struct {
	ID         int64     `db:"id"`
	EmployerID int64     `db:"employer_id"`
	Content    string    `db:"content"`
	IsFree     bool      `db:"isfree"`
	Amount     float32   `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
}

type SubscriptionReq struct {
	JobseekerID  int64 `db:"jobseeker_id"`
	NewsLetterID int64 `db:"newletter_id"`
}

type PaymentReq struct {
	SubscriptionID int64   `db:"subscription_id"`
	Amount         float32 `db:"amount"`
}

type RazorpayReq struct {
	PaymentID int64  `db:"payment_id"`
	OrderID   string `db:"order_id"`
}

type SubscriptionRes struct {
	ID          int64        `db:"id"`
	JobSeekerID int64        `db:"jobseeker_id"`
	NewLetterID int64        `db:"newletter_id"`
	StartDate   sql.NullTime `db:"startdate"`
	EndDate     sql.NullTime `db:"enddate"`
	Status      string       `db:"status"`
}

type PaymentRes struct {
	ID             int64     `db:"id"`
	SubscriptionID int64     `db:"subscription_id"`
	Amount         float32   `db:"amount"`
	Status         string    `db:"status"`
	Date           time.Time `db:"date"`
}

type RazorpayRes struct {
	ID        int64          `db:"id"`
	PaymentID int64          `db:"payment_id"`
	PayID     sql.NullString `db:"pay_id"`
	OrderID   sql.NullString `db:"order_id"`
	Signature sql.NullString `db:"signature"`
}

type Razorpay struct {
	PayID     string `db:"pay_id"`
	OrderID   string `db:"order_id"`
	Signature string `db:"signature"`
}

type SubscriptionDate struct {
	StartDate sql.NullTime `db:"startdate"`
	EndDate   sql.NullTime `db:"enddate"`
}
