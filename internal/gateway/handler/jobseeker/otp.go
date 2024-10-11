package jobseeker

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	config "github.com/ahdaan67/jobportal/config"
)

var otpMap = make(map[string]string)

func GenerateOTP() string {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	return fmt.Sprintf("%06d", randGen.Intn(1000000))
}

func SendOTP(email string, cfg config.Config) (string, error) {
	otp := GenerateOTP()
	from := cfg.Email
	password := cfg.Password
	to := email
	log.Println("Sending OTP to email:", email, "OTP:", otp)
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	otpMap[email] = otp

	auth := smtp.PlainAuth("", from, password, smtpServer)
	message := fmt.Sprintf("Subject: Your OTP\n\nYour OTP is: %s", otp)

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println("Error sending OTP:", err)
		return "", err
	}

	return otp, nil
}

func VerifyOTP(otp string, email string) bool {
	storedOTP, ok := otpMap[email]
	if !ok {
		log.Println("OTP not found for email:", email)
		return false
	}
	if otp == storedOTP {
		delete(otpMap, email)
		return true
	} else {
		return false
	}
}