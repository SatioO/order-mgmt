package mapper

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/satioO/order-mgmt/pkg/models"
	"github.com/satioO/order-mgmt/pkg/pb"
)

func ToCreateOrderEntity(dto *pb.CreateOrderRequest) models.Order {
	uuid, _ := uuid.NewUUID()

	return models.Order{
		PK:               fmt.Sprintf("ORDER#%s", uuid.String()),
		SK:               fmt.Sprintf("ORDER#%s", uuid.String()),
		Type:             "Order",
		CustomerID:       dto.CustomerId,
		SellerID:         dto.SellerId,
		PaymentMethod:    dto.PaymentMethod.String(),
		DeliveryLocation: dto.DeliveryLocation,
		OrderStatus:      "ORDER_CREATED",
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}
}
