package main

import (
	"net"

	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/db"
	"github.com/ahdaan67/jobportal/internal/newsletter/api"
	"github.com/ahdaan67/jobportal/internal/newsletter/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/newsletter"
	"google.golang.org/grpc"
)

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("/root/logs/Newsletter.log")
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

	nt := storer.NewNEWSLETTERStorer(db.GetDB())
	nrv := api.NewServer(nt, *cfg)

	grpcSrv := grpc.NewServer()
	pb.RegisterNewsLetterServer(grpcSrv, nrv)
	logrusLogger.Info("Newsletter server registered successfully.")

	listener, err := net.Listen("tcp", cfg.NewsLetterPort)
	if err != nil {
		logrusLogger.Fatalf("listener failed: %v", err)
	}
	logrusLogger.Infof("Server listening on %s", cfg.NewsLetterPort)

	err = grpcSrv.Serve(listener)
	if err != nil {
		logrusLogger.Fatalf("failed to serve: %v", err)
	}
	logrusLogger.Info("Server stopped gracefully.")
}