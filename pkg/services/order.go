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
	repo *models.OrderRepo
}

func NewOrderService(repo *models.OrderRepo) *orderService {
	return &orderService{
		repo: repo,
	}
}

func (o orderService) CreateOrder(ctx context.Context, body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	orderId, _ := uuid.NewUUID()

	order := models.Order{
		PK:         fmt.Sprintf("ORDER#%s", orderId),
		SK:         fmt.Sprintf("ORDER#%s", orderId),
		CustomerID: body.CustomerId,
		SellerID:   body.SellerId,
		Product: models.Product{
			ProductID:   body.Product.ProductId,
			ProductName: body.Product.ProductName,
		},
		PaymentMethod:    body.PaymentMethod.String(),
		DeliveryLocation: body.DeliveryLocation,
		OrderStatus:      "ORDER_CREATED",
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}

	return o.repo.CreateOrder(order)
}
