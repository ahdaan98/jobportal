package main

import (
	"log"
	"net"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/db"
	"github.com/ahdaan67/jobportal/internal/newsletter/api"
	"github.com/ahdaan67/jobportal/internal/newsletter/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/newsletter"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	defer db.Close()
	log.Println("successfully connected to database")

	nt := storer.NewNEWSLETTERStorer(db.GetDB())
	nrv := api.NewServer(nt, *cfg)

	grpcSrv := grpc.NewServer()
	pb.RegisterNewsLetterServer(grpcSrv, nrv)

	listener, err := net.Listen("tcp", cfg.NewsLetterPort)
	if err != nil {
		log.Fatalf("listener failed: %v", err)
	}

	log.Printf("server listening on %s", cfg.NewsLetterPort)
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
