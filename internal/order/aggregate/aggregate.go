package aggregate

import (
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
	return nil
}
