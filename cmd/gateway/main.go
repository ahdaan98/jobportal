package main

import (
	"log"

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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	jobClient, err := CreateJobClient(opts, *cfg)
	if err != nil {
		log.Fatalf("failed to connect to job service: %v", err)
	}

	jobseekerClient, employerClient, err := CreateUserClient(opts, *cfg)
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	newsletterClient, err := CreateNewsLetterClient(opts, *cfg)
	if err != nil {
		log.Fatalf("failed to connect to newsletter service: %v", err)
	}

	jsh := jobseeker.NewHandler(jobseekerClient, *cfg)
	emh := employer.NewHandler(employerClient, *cfg)
	jhl := job.NewHandler(jobClient, *cfg)
	nhl := newsletter.NewHandler(newsletterClient, *cfg)
	vhl := handler.NewVideoCallHandler()

	handler.RegisterRoutes(jhl, jsh, emh, vhl, nhl)

	if err := handler.Start(cfg.GatewayPort); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func CreateJobClient(opts []grpc.DialOption, cfg config.Config) (pb.JobClient, error) {
	conn, err := grpc.NewClient(cfg.JobPort, opts...)
	if err != nil {
		return nil, err
	}

	return pb.NewJobClient(conn), nil
}

func CreateUserClient(opts []grpc.DialOption, cfg config.Config) (js.JobSeekerClient, em.EmployerClient, error) {
	conn, err := grpc.NewClient(cfg.UserPort, opts...)
	if err != nil {
		return nil, nil, err
	}

	return js.NewJobSeekerClient(conn), em.NewEmployerClient(conn), nil
}

func CreateNewsLetterClient(opts []grpc.DialOption, cfg config.Config) (nl.NewsLetterClient, error) {
	conn, err := grpc.NewClient(cfg.NewsLetterPort, opts...)
	if err != nil {
		return nil, err
	}

	return nl.NewNewsLetterClient(conn), nil
}
