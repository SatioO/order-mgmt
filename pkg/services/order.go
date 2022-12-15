package services

import (
	"context"

	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
)

type orderService struct {
	pb.UnimplementedOrderServiceServer
	repo *models.OrderRepo
}

func NewOrderService(repo *models.OrderRepo) *orderService {
	return &orderService{
		repo: repo,
	}
}

func (o orderService) CreateOrder(ctx context.Context, body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	o.repo.CreateOrder(body)
	return &pb.OrderResponse{}, nil
}
