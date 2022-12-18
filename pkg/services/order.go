package services

import (
	"context"
	"sync"

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

	statusChan := make(chan error)
	go o.CreateOrderItem(orderEntity.PK, body.Products, statusChan)

	for err := range statusChan {
		if err != nil {
			return nil, err
		}
	}

	return createdOrder, err
}

func (o orderService) CreateOrderItem(orderId string, products []*pb.Product, statusChan chan<- error) {
	defer close(statusChan)

	wg := sync.WaitGroup{}
	wg.Add(len(products))

	for _, product := range products {
		go func(product *pb.Product) {
			orderItemEntity := mapper.ToCreateOrderItemEntity(orderId, product)
			statusChan <- o.orderItemRepo.CreateOrderItem(orderItemEntity)

			wg.Done()
		}(product)
	}

	wg.Wait()
}
