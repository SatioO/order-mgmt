package services

import (
	"context"
	"fmt"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(body.Products))

	for _, product := range body.Products {
		go func(product *pb.Product) {
			orderItemId, _ := uuid.NewUUID()
			o.orderItemRepo.CreateOrderItem(models.OrderItem{
				PK:               fmt.Sprintf("ORDER#%s", orderId),
				SK:               fmt.Sprintf("ORDERITEM#%s", orderItemId),
				ProductName:      product.ProductName,
				CreatedTimestamp: time.Now().String(),
				UpdatedTimestamp: time.Now().String(),
			})
			wg.Done()
		}(product)
	}

	wg.Wait()

	return createdOrder, err
}
