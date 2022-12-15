package services

import (
	"context"

	"github.com/satioO/order-mgmt/pkg/pb"
)

type orderService struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderService() *orderService {
	return &orderService{}
}

func (orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}
