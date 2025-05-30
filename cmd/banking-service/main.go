package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LavaJover/shvark-banking-service/internal/config"
	grpcapi "github.com/LavaJover/shvark-banking-service/internal/delivery/grpc"
	"github.com/LavaJover/shvark-banking-service/internal/infrastructure/postgres"
	"github.com/LavaJover/shvark-banking-service/internal/usecase"
	bankingpb "github.com/LavaJover/shvark-banking-service/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	// Reading config
	cfg := config.MustLoad()

	// Init database
	db := postgres.MustInitDB(cfg)

	// Init user repo
	bankDetailRepo, err := postgres.NewDefaultbankDetailRepository(db)
	if err != nil{
		log.Fatalf("failed to init user repository: %v\n", err.Error())
	}

	// Init user usecase
	uc := usecase.NewDefaultBankDetailUsecase(bankDetailRepo)

	// Creating gRPC server
	grpcServer := grpc.NewServer()
	bankDetailHandler := grpcapi.NewbankDetailHandler(uc)

	bankingpb.RegisterBankingServiceServer(grpcServer, bankDetailHandler)

	// Start
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server started on %s:%s\n", cfg.Host, cfg.Port)
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v\n", err)
	}
}