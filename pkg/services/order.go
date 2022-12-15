package services

import (
	"context"

	"github.com/satioO/order-mgmt/pkg/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func (OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	return nil, nil
}
