package mapper

import (
	"time"

	"github.com/satioO/order-mgmt/internal/models"
	"github.com/satioO/order-mgmt/proto"
	"github.com/segmentio/ksuid"
)

func ToCreateOrderItemEntity(orderId string, orderItem *proto.Product) models.OrderItem {
	orderItemId := ksuid.New().String()

	return models.OrderItem{
		PK:               orderId + "#ITEM#" + orderItemId,
		SK:               orderId + "#ITEM#" + orderItemId,
		GSI1PK:           orderId,
		GSI1SK:           "ITEM#" + orderItemId,
		Type:             "OrderItem",
		ProductName:      orderItem.ProductName,
		CreatedTimestamp: time.Now().UTC().String(),
		UpdatedTimestamp: time.Now().UTC().String(),
	}
}
