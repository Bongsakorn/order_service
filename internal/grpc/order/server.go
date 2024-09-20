package grpcorder

import (
	"context"
	"log"
	pb "order_service/internal/grpc/protobuf"
	"order_service/internal/repository"
	repositoryModel "order_service/internal/repository/models"
)

// Server struct
type Server struct {
	pb.UnimplementedOrdersServiceServer
	Repository *repository.Repository
}

// GetOrder returns a list of orders.
func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	log.Printf("Geting order %d...\n", req.GetId())
	var order repositoryModel.Order
	var resp *pb.Order
	s.Repository.Db.Find(&order, req.GetId())

	resp = &pb.Order{
		Id:         order.ID,
		Status:     order.Status,
		RentalType: order.RentalType,
		PickupDate: order.PickupDate.Unix(),
		ReturnDate: order.ReturnDate.Unix(),
		TotalPrice: order.TotalPrice,
		RentalDays: order.RentalDays,
		CreatedAt:  order.CreatedAt.Unix(),
		UpdatedAt:  order.UpdatedAt.Unix(),
	}

	log.Println(resp)

	return resp, nil
}
