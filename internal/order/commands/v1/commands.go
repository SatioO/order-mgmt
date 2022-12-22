package v1

import (
	"github.com/satioO/order-mgmt/internal/order/models"
	"github.com/satioO/order-mgmt/pkg/es"
)

type CreateOrderCommand struct {
	es.BaseCommand
	CustomerID       string              `json:"customerId" validate:"required"`
	PaymentMethod    string              `json:"paymentMethod" validdate:"required"`
	DeliveryLocation string              `json:"deliveryLocation" validate:"required"`
	OrderItems       []*models.OrderItem `json:"items" validate:"required"`
}

func NewCreateOrderCommand(aggregateID string, customerId string, paymentMethod string, deliveryLocation string, orderItems []*models.OrderItem) *CreateOrderCommand {
	return &CreateOrderCommand{BaseCommand: *es.NewBaseCommand(aggregateID), CustomerID: customerId, PaymentMethod: paymentMethod, DeliveryLocation: deliveryLocation, OrderItems: orderItems}
}
