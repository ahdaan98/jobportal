package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Email    string `yaml:"EMAIL"`
	Password string `yaml:"PASSWORD"`

	LinkedinClientID       string `yaml:"LINKEDIN_CLIENT_ID"`
	LinkedinClientSecretID string `yaml:"LINKEDIN_CLIENT_SECRET"`

	AcessKeyID      string `yaml:"ACCESSKEYID"`
	SecretAccessKey string `yaml:"SECRETACCESSKEY"`
	AwsRegion       string `yaml:"AWSREGION"`

	RazorpayKey    string `yaml:"RAZORPAYKEY"`
	RazorpaySecret string `yaml:"RAZORPAYSECRET"`

	GatewayPort    string `yaml:"GATEWAYPORT"`
	JobPort        string `yaml:"JOBPORT"`
	UserPort       string `yaml:"USERPORT"`
	NewsLetterPort string `yaml:"NEWSLETTERPORT"`
}

func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %w", err)
	}

	return &config, nil
}
