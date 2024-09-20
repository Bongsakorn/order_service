package grpc

import (
	"fmt"
	"log"
	"net"
	"order_service/internal/config"
	grpcorder "order_service/internal/grpc/order"
	"order_service/internal/repository"

	pb "order_service/internal/grpc/protobuf"

	"google.golang.org/grpc"
)

// StartGrpcServer starts the gRPC server.
func StartGrpcServer(config *config.ConfigService) {
	log.Println("Starting grpc server...")
	repository := repository.NewRepository(config)
	server := &grpcorder.Server{
		Repository: repository,
	}

	s := grpc.NewServer()
	pb.RegisterOrdersServiceServer(s, server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Println("GRPC Server start on port", config.GrpcPort)
	}
}
