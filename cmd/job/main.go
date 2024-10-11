package main

import (
	"log"
	"net"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/db"
	"github.com/ahdaan67/jobportal/internal/job/api"
	"github.com/ahdaan67/jobportal/internal/job/storer"
	pb "github.com/ahdaan67/jobportal/utils/pb/job"
	"google.golang.org/grpc"
)

func main() {
	cfg,err := config.LoadConfig("config/config.yaml")
	if err!=nil {
		log.Fatalf("cannot load config: %v",err)
	}
	
	db, err := db.NewDatabase()
	if err!=nil{
		log.Fatalf("error opening database: %v", err)
	}

	defer db.Close()
	log.Println("successfully connected to database")

	jt := storer.NewJOBStorer(db.GetDB())
	jrv := api.NewServer(jt, *cfg)

	grpcSrv := grpc.NewServer()
	pb.RegisterJobServer(grpcSrv, jrv)

	listener, err := net.Listen("tcp", cfg.JobPort)
	if err != nil {
		log.Fatalf("listener failed: %v", err)
	}

	log.Printf("server listening on %s", cfg.JobPort)
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}