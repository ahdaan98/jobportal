package main

import (
	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/internal/gateway/handler"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/employer"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/job"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/jobseeker"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/newsletter"
	em "github.com/ahdaan67/jobportal/utils/pb/employer"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	js "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	nl "github.com/ahdaan67/jobportal/utils/pb/newsletter"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const LogPath = "/root/logs/Gateway.log"

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger(LogPath)
	defer logrusLogFile.Close()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logrusLogger.Fatalf("Cannot load config: %v", err)
	}
	logrusLogger.Info("Configuration loaded successfully.")

	jobClient, err := CreateJobClient(opts, *cfg)
	if err != nil {
		logrusLogger.Fatalf("Failed to connect to job service: %v", err)
	}
	logrusLogger.Info("Connected to job service.")

	jobseekerClient, employerClient, err := CreateUserClient(opts, *cfg)
	if err != nil {
		logrusLogger.Fatalf("Failed to connect to user service: %v", err)
	}
	logrusLogger.Info("Connected to user service.")

	newsletterClient, err := CreateNewsLetterClient(opts, *cfg)
	if err != nil {
		logrusLogger.Fatalf("Failed to connect to newsletter service: %v", err)
	}
	logrusLogger.Info("Connected to newsletter service.")

	jsh := jobseeker.NewHandler(jobseekerClient, *cfg, LogPath)
	emh := employer.NewHandler(employerClient, *cfg, LogPath)
	jhl := job.NewHandler(jobClient, *cfg, LogPath)
	nhl := newsletter.NewHandler(newsletterClient, jobseekerClient, *cfg)
	vhl := handler.NewVideoCallHandler()

	handler.RegisterRoutes(jhl, jsh, emh, vhl, nhl)
	logrusLogger.Info("Routes registered successfully.")

	if err = handler.Start(cfg.GatewayPort); err != nil {
		logrusLogger.Fatalf("Failed to run server: %v", err)
	}
	logrusLogger.Infof("Server started and listening on port: %s", cfg.GatewayPort)
}

func CreateJobClient(opts []grpc.DialOption, cfg config.Config) (pb.JobClient, error) {
	conn, err := grpc.NewClient(cfg.JobPort, opts...)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Job client created successfully.")
	return pb.NewJobClient(conn), nil
}

func CreateUserClient(opts []grpc.DialOption, cfg config.Config) (js.JobSeekerClient, em.EmployerClient, error) {
	conn, err := grpc.NewClient(cfg.UserPort, opts...)
	if err != nil {
		return nil, nil, err
	}
	logrus.Infof("User client created successfully.")
	return js.NewJobSeekerClient(conn), em.NewEmployerClient(conn), nil
}

func CreateNewsLetterClient(opts []grpc.DialOption, cfg config.Config) (nl.NewsLetterClient, error) {
	conn, err := grpc.NewClient(cfg.NewsLetterPort, opts...)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Newsletter client created successfully.")
	return nl.NewNewsLetterClient(conn), nil
}
