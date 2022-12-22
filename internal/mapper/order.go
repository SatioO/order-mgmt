package mapper

import (
	"time"

	"github.com/satioO/order-mgmt/internal/models"
	"github.com/satioO/order-mgmt/proto"
	"github.com/segmentio/ksuid"
)

func ToCreateOrderEntity(dto *proto.CreateOrderRequest) models.Order {
	orderId := ksuid.New().String()

	return models.Order{
		PK:               "CUSTOMER#" + dto.CustomerId,
		SK:               "ORDER#" + orderId,
		GSI1PK:           "ORDER#" + orderId,
		GSI1SK:           "ORDER#" + orderId,
		Type:             "Order",
		CustomerID:       dto.CustomerId,
		SellerID:         dto.SellerId,
		OrderID:          orderId,
		PaymentMethod:    dto.PaymentMethod.String(),
		DeliveryLocation: dto.DeliveryLocation,
		OrderStatus:      "ORDER_CREATED",
		CreatedTimestamp: time.Now().UTC().String(),
		UpdatedTimestamp: time.Now().UTC().String(),
	}
}
