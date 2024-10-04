package service

import (
	"bytes"
	"fmt"

	"github.com/ahdaan67/jobportal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func AddImageToAwsS3(cfg config.Config,file []byte, filename string) (string, error) {

	fmt.Println("print1", cfg.AwsRegion)
	fmt.Println("print2", cfg.AcessKeyID)
	fmt.Println("print3", cfg.SecretAccessKey)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AwsRegion),
		Credentials: credentials.NewStaticCredentials(
			cfg.AcessKeyID,
			cfg.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		fmt.Println("erorrrr here", err)
		return "", err
	}

	uploader := s3manager.NewUploader(sess)
	bucketName := "jobquestbucket"

	fmt.Printf("Uploading file of size: %d bytes\n", len(file))

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("application/pdf"),
	})

	if err != nil {
		fmt.Println("erroorrrr 2", err)
		return "", err
	}
	fmt.Println("Bucket: ", bucketName)
	fmt.Println("aws region: ", cfg.AwsRegion)
	fmt.Println("file name: ", filename)
	fmt.Printf("Upload result: %+v\n", result)
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)
	return url, nil
}