package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
)

func ToCreateOrderItemEntity(orderId string, orderItem *pb.Product) models.OrderItem {
	orderItemId, _ := uuid.NewUUID()

	return models.OrderItem{
		PK:               orderId,
		SK:               "ORDERITEM#" + orderItemId.String(),
		Type:             "OrderItem",
		ProductName:      orderItem.ProductName,
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}
}
