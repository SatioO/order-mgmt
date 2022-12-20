package mapper

import (
	"time"

	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/segmentio/ksuid"
)

func ToCreateOrderEntity(dto *pb.CreateOrderRequest) models.Order {
	orderId := ksuid.New()

	return models.Order{
		PK:               "ORDER#" + orderId.String(),
		SK:               "METADATA#" + orderId.String(),
		GSI1PK:           "ORDER",
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
