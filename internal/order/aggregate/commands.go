package aggregate

import (
	"context"

	"github.com/pkg/errors"
	v1 "github.com/satioO/order-mgmt/internal/order/events/v1"
	"github.com/satioO/order-mgmt/internal/order/models"
)

func (a *OrderAggregate) CreateOrder(ctx context.Context, orderItems []*models.OrderItem, paymentMethod, deliveryLocation string) error {
	event, err := v1.NewOrderCreatedEvent(a, orderItems, paymentMethod, deliveryLocation)
	if err != nil {
		return errors.Wrap(err, "NewOrderCreatedEvent")
	}

	return a.Apply(event)
}
