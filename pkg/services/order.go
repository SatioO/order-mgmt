package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	orderId, _ := uuid.NewUUID()
	executed := make(chan error)

	order := models.Order{
		PK:               fmt.Sprintf("ORDER#%s", orderId),
		SK:               fmt.Sprintf("ORDER#%s", orderId),
		CustomerID:       body.CustomerId,
		SellerID:         body.SellerId,
		PaymentMethod:    body.PaymentMethod.String(),
		DeliveryLocation: body.DeliveryLocation,
		OrderStatus:      "ORDER_CREATED",
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}

	createdOrder, err := o.orderRepo.CreateOrder(order)

	go o.CreateOrderItem(orderId, body.Products, executed)

	for err := range executed {
		if err != nil {
			return nil, err
		}
	}

	return createdOrder, err
}

func (o orderService) CreateOrderItem(orderId uuid.UUID, products []*pb.Product, executed chan<- error) {
	defer close(executed)

	for _, product := range products {
		orderItemId, _ := uuid.NewUUID()
		err := o.orderItemRepo.CreateOrderItem(models.OrderItem{
			PK:               fmt.Sprintf("ORDER#%s", orderId),
			SK:               fmt.Sprintf("ORDERITEM#%s", orderItemId),
			ProductName:      product.ProductName,
			CreatedTimestamp: time.Now().String(),
			UpdatedTimestamp: time.Now().String(),
		})

		executed <- err
	}
}
