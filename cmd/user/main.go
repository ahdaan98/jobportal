package main

import (
	"log"
	"net"

	"github.com/ahdaan67/jobportal/config"
	"github.com/ahdaan67/jobportal/db"
	api "github.com/ahdaan67/jobportal/internal/user/api/jobseeker"
	epi "github.com/ahdaan67/jobportal/internal/user/api/employer"
	estorer "github.com/ahdaan67/jobportal/internal/user/storer/employer"
	storer "github.com/ahdaan67/jobportal/internal/user/storer/jobseeker"
	pb "github.com/ahdaan67/jobportal/utils/pb/jobseeker"
	eb "github.com/ahdaan67/jobportal/utils/pb/employer"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	
	db, err := db.NewDatabase()
	if err!=nil{
		log.Fatalf("error opening database: %v", err)
	}

	defer db.Close()
	log.Println("successfully connected to database")

	jt := storer.NewJOBSEEKERStorer(db.GetDB())
	et := estorer.NewEMPLOYERstorer(db.GetDB())

	jrv := api.NewServer(jt, et)
	erv := epi.NewServer(et)


	grpcSrv := grpc.NewServer()
	pb.RegisterJobSeekerServer(grpcSrv, jrv)
	eb.RegisterEmployerServer(grpcSrv, erv)

	listener, err := net.Listen("tcp", cfg.UserPort)
	if err != nil {
		log.Fatalf("listener failed: %v", err)
	}

	log.Printf("server listening on %s", cfg.UserPort)
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}