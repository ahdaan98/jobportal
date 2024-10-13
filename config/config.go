package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email              string
	Password           string
	LinkedinClientID   string
	LinkedinClientSecretID string
	AcessKeyID        string
	SecretAccessKey   string
	AwsRegion         string
	RazorpayKey       string
	RazorpaySecret    string
	GatewayPort       string
	JobPort           string
	UserPort          string
	NewsLetterPort    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &Config{
		Email:              getEnv("EMAIL", ""),
		Password:           getEnv("PASSWORD", ""),
		LinkedinClientID:   getEnv("LINKEDIN_CLIENT_ID", ""),
		LinkedinClientSecretID: getEnv("LINKEDIN_CLIENT_SECRET", ""),
		AcessKeyID:        getEnv("ACCESSKEYID", ""),
		SecretAccessKey:   getEnv("SECRETACCESSKEY", ""),
		AwsRegion:         getEnv("AWSREGION", ""),
		RazorpayKey:       getEnv("RAZORPAYKEY", ""),
		RazorpaySecret:    getEnv("RAZORPAYSECRET", ""),
		GatewayPort:       getEnv("GATEWAYPORT", ""),
		JobPort:           getEnv("JOBPORT", ""),
		UserPort:          getEnv("USERPORT", ""),
		NewsLetterPort:    getEnv("NEWSLETTERPORT", ""),
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}