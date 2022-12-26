package delivery

import (
	"context"

	v1 "github.com/satioO/order-mgmt/internal/order/commands/v1"
	"github.com/satioO/order-mgmt/internal/order/models"
	"github.com/satioO/order-mgmt/internal/order/service"
	"github.com/satioO/order-mgmt/pkg/logger"
	"github.com/satioO/order-mgmt/proto"
	"github.com/segmentio/ksuid"
)

type orderGRPCService struct {
	proto.UnimplementedOrderServiceServer
	log logger.Logger
	os  *service.OrderService
}

func NewOrderGRPCService(log logger.Logger, os *service.OrderService) *orderGRPCService {
	return &orderGRPCService{os: os, log: log}
}

func (o *orderGRPCService) GetOrders(ctx context.Context, req *proto.GetOrdersRequest) (*proto.GetOrdersResponse, error) {
	return nil, nil
}

func (o *orderGRPCService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	aggregateID := ksuid.New().String()
	command := v1.NewCreateOrderCommand(aggregateID, req.CustomerId, req.PaymentMethod.String(), req.DeliveryLocation, []*models.OrderItem{})

	if err := o.os.Commands.CreateOrder.Handle(ctx, command); err != nil {
		return nil, err
	}

	return &proto.OrderResponse{OrderId: aggregateID}, nil
}
