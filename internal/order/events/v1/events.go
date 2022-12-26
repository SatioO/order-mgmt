package v1

import (
	"github.com/satioO/order-mgmt/internal/order/models"
	"github.com/satioO/order-mgmt/pkg/es"
)

const (
	OrderCreated = "V1_ORDER_CREATED"
)

type OrderCreatedEvent struct {
	OrderItems       []*models.OrderItem `json:"orderItems"`
	PaymentMethod    string              `json:"paymentMethod"`
	DeliveryLocation string              `json:"deliveryLocation"`
}

func NewOrderCreatedEvent(aggregate es.Aggregate, orderItems []*models.OrderItem, paymentMethod string, deliveryLocation string) (es.Event, error) {
	eventData := &OrderCreatedEvent{
		OrderItems:       orderItems,
		PaymentMethod:    paymentMethod,
		DeliveryLocation: deliveryLocation,
	}

	event := es.NewBaseEvent(aggregate, OrderCreated)

	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}

	return event, nil
}
