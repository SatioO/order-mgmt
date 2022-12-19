package services

import (
	"context"

	"github.com/satioO/order-mgmt/pkg/mapper"
	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
)

type orderService struct {
	pb.UnimplementedOrderServiceServer
	orderRepo     *models.OrderRepo
	orderItemRepo *models.OrderItemRepo
}

func NewOrderService(orderRepo *models.OrderRepo, orderItemRepo *models.OrderItemRepo) *orderService {
	return &orderService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (o orderService) GetOrders(ctx context.Context, body *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	return o.orderRepo.GetOrders(ctx)
}

func (o orderService) CreateOrder(ctx context.Context, body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	orderEntity := mapper.ToCreateOrderEntity(body)
	createdOrder, err := o.orderRepo.CreateOrder(ctx, orderEntity)

	var orderItems []models.OrderItem
	for _, product := range body.Products {
		orderItemEntity := mapper.ToCreateOrderItemEntity(orderEntity.PK, product)
		orderItems = append(orderItems, orderItemEntity)
	}

	if err := o.orderItemRepo.BatchWriteOrderItems(orderItems); err != nil {
		return nil, err
	}

	return createdOrder, err
}
