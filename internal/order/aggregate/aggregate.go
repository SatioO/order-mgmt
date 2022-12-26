package aggregate

import (
	"github.com/pkg/errors"
	v1 "github.com/satioO/order-mgmt/internal/order/events/v1"
	"github.com/satioO/order-mgmt/internal/order/models"
	"github.com/satioO/order-mgmt/pkg/es"
)

const (
	OrderAggregateType es.AggregateType = "order"
)

type OrderAggregate struct {
	*es.AggregateBase
	Order *models.Order
}

func NewOrderAggregateWithID(id string) *OrderAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewOrderAggregate()
	aggregate.SetID(id)
	aggregate.Order.ID = id
	return aggregate
}

func NewOrderAggregate() *OrderAggregate {
	orderAggregate := &OrderAggregate{Order: models.NewOrder()}
	base := es.NewAggregateBase(orderAggregate.When)
	base.SetType(OrderAggregateType)
	orderAggregate.AggregateBase = base
	return orderAggregate
}

func (a *OrderAggregate) When(evt es.Event) error {
	switch evt.GetEventType() {
	case v1.OrderCreated:
		return a.onOrderCreated(evt)
	default:
		return es.ErrInvalidEventType
	}
}

func (a *OrderAggregate) onOrderCreated(evt es.Event) error {
	var eventData v1.OrderCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Order.DeliveryLocation = eventData.DeliveryLocation
	a.Order.PaymentMethod = eventData.PaymentMethod
	return nil
}
