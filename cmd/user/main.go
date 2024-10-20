package main

import (
	"net"

	logging "github.com/ahdaan67/jobportal/logging"
	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/db"
	epi "github.com/ahdaan67/jobportal/internal/user/api/employer"
	api "github.com/ahdaan67/jobportal/internal/user/api/jobseeker"
	estorer "github.com/ahdaan67/jobportal/internal/user/storer/employer"
	storer "github.com/ahdaan67/jobportal/internal/user/storer/jobseeker"
	eb "github.com/ahdaan67/jobportal/utils/pb/employer"
	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	"google.golang.org/grpc"
)

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("/root/logs/User.log")
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

	jt := storer.NewJOBSEEKERStorer(db.GetDB())
	et := estorer.NewEMPLOYERstorer(db.GetDB())

	jrv := api.NewServer(jt, et)
	erv := epi.NewServer(et)

	grpcSrv := grpc.NewServer()
	pb.RegisterJobSeekerServer(grpcSrv, jrv)
	eb.RegisterEmployerServer(grpcSrv, erv)
	logrusLogger.Info("JobSeeker and Employer servers registered successfully.")

	listener, err := net.Listen("tcp", cfg.UserPort)
	if err != nil {
		logrusLogger.Fatalf("listener failed: %v", err)
	}
	logrusLogger.Infof("Server listening on %s", cfg.UserPort)

	err = grpcSrv.Serve(listener)
	if err != nil {
		logrusLogger.Fatalf("failed to serve: %v", err)
	}
	logrusLogger.Info("Server stopped gracefully.")
}
