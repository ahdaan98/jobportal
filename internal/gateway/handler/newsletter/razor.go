package newsletter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

func CreateRazorpayOrder(client *razorpay.Client, amount float64) (string, error) {
	fmt.Println("order amount", amount)

	data := map[string]interface{}{
		"amount":   int(amount * 100),
		"currency": "INR",
		"receipt":  "some_receipt_id" + strconv.Itoa(int(amount)),
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", err
	}

	return body["id"].(string), nil
}

func VerifyRazorpaySignature(orderId, paymentId, razorpaySignature, secret string) error {
	message := orderId + "|" + paymentId
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	if expectedSignature != razorpaySignature {
		return errors.New("signature verification failed")
	}
	return nil
}