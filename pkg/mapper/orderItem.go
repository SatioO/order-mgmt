package mapper

import (
	"time"

	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/segmentio/ksuid"
)

func ToCreateOrderItemEntity(orderId string, orderItem *pb.Product) models.OrderItem {
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
