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

func (o orderService) CreateOrder(ctx context.Context, body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	orderEntity := mapper.ToCreateOrderEntity(body)
	createdOrder, err := o.orderRepo.CreateOrder(orderEntity)

	var orderItems []models.OrderItem
	for _, items := range body.Products {
		orderItemEntity := mapper.ToCreateOrderItemEntity(orderEntity.PK, items)
		orderItems = append(orderItems, orderItemEntity)
	}

	if err := o.orderItemRepo.BatchWriteOrderItems(orderItems); err != nil {
		return nil, err
	}

	return createdOrder, err
}
