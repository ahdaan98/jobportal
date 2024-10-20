package main

import (
	"net"

	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/db"
	"github.com/ahdaan67/jobportal/internal/job/api"
	"github.com/ahdaan67/jobportal/internal/job/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	"google.golang.org/grpc"
)

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("/root/logs/Job.log")
	defer logrusLogFile.Close()

	cfg, err := config.LoadConfig()
	if err != nil {
		logrusLogger.Fatalf("cannot load config: %v", err)
	}
	logrusLogger.Info("Configuration loaded successfully.")

	db, err := db.NewDatabase()
	if err != nil {
		logrusLogger.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	logrusLogger.Info("Successfully connected to database.")

	jt := storer.NewJOBStorer(db.GetDB())
	jrv := api.NewServer(jt, *cfg)

	grpcSrv := grpc.NewServer()
	pb.RegisterJobServer(grpcSrv, jrv)
	logrusLogger.Info("Job server registered successfully.")

	listener, err := net.Listen("tcp", cfg.JobPort)
	if err != nil {
		logrusLogger.Fatalf("listener failed: %v", err)
	}
	logrusLogger.Infof("Server listening on %s", cfg.JobPort)

	err = grpcSrv.Serve(listener)
	if err != nil {
		logrusLogger.Fatalf("failed to serve: %v", err)
	}
	logrusLogger.Info("Server stopped gracefully.")
}