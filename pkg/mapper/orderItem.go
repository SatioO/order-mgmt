package mapper

import (
	"time"

	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/segmentio/ksuid"
)

func ToCreateOrderItemEntity(orderId string, orderItem *pb.Product) models.OrderItem {
	orderItemId := ksuid.New()

	return models.OrderItem{
		PK:               orderId + "#ITEM#" + orderItemId.String(),
		SK:               orderId + "#ITEM#" + orderItemId.String(),
		GSI1PK:           orderId,
		GSI1SK:           "ITEM#" + orderItemId.String(),
		Type:             "OrderItem",
		ProductName:      orderItem.ProductName,
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}
}
