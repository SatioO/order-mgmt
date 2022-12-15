package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/satioO/order-mgmt/pkg/pb"
)

type orderService struct {
	pb.UnimplementedOrderServiceServer
	db *dynamodb.Client
}

func NewOrderService(db *dynamodb.Client) *orderService {
	return &orderService{
		db: db,
	}
}

func (orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}
