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
		SK:               "ORDER#" + orderId.String(),
		GSI1PK:           "Order",
		GSI1SK:           "ORDER#" + orderId.String(),
		Type:             "Order",
		CustomerID:       dto.CustomerId,
		SellerID:         dto.SellerId,
		OrderID:          orderId.String(),
		PaymentMethod:    dto.PaymentMethod.String(),
		DeliveryLocation: dto.DeliveryLocation,
		OrderStatus:      "ORDER_CREATED",
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}
}
