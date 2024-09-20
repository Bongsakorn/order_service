package grpc

import (
	"fmt"
	"log"
	"net"
	"order_service/internal/config"
	"order_service/internal/repository"

	"google.golang.org/grpc"

	grpcexample "order_service/internal/grpc/grpc-example"
	grpcexample2 "order_service/internal/grpc/grpc-example2"
	pb "order_service/internal/grpc/protobuf"
)

func StartGrpcServer(config *config.ConfigService) {
	repository := repository.NewRepository(config)
	server := &grpcexample.Server{
		Repository: repository,
	}
	server2 := &grpcexample2.Server{}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, server)
	pb.RegisterGreetingServiceServer(s, server2)
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
