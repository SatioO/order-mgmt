package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
)

func ToCreateOrderEntity(dto *pb.CreateOrderRequest) models.Order {
	orderId, _ := uuid.NewUUID()

	return models.Order{
		PK:               "ORDER#" + orderId.String(),
		SK:               "METADATA#" + orderId.String(),
		GSI1PK:           "ORDER",
		GSI1SK:           orderId.String(),
		Type:             "Order",
		CustomerID:       dto.CustomerId,
		SellerID:         dto.SellerId,
		OrderID:          orderId.String(),
		PaymentMethod:    dto.PaymentMethod.String(),
		DeliveryLocation: dto.DeliveryLocation,
		OrderStatus:      "CREATED",
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}
}
