package mapper

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
)

func ToCreateOrderItemEntity(orderId string, orderItem *pb.Product) models.OrderItem {
	orderItemId, _ := uuid.NewUUID()

	return models.OrderItem{
		PK:               orderId,
		SK:               fmt.Sprintf("ORDERITEM#%s", orderItemId),
		Type:             "OrderItem",
		ProductName:      orderItem.ProductName,
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}
}
