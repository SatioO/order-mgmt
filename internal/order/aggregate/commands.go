package aggregate

import (
	"context"

	"github.com/satioO/order-mgmt/internal/order/models"
)

func (a *OrderAggregate) CreateOrder(ctx context.Context, orderItems []*models.OrderItem, paymentMethod, deliveryLocation string) error {
	return nil
}
