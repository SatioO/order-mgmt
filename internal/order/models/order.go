package models

import "fmt"

type Order struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	GSI1PK           string `dynamodbav:"GSI1-PK"`
	GSI1SK           string `dynamodbav:"GSI1-SK"`
	ID               string `dynamodbav:"id"`
	Type             string `dynamodbav:"type"`
	CustomerID       string `dynamodbav:"customerId"`
	PaymentMethod    string `dynamodbav:"paymentMethod"`
	DeliveryLocation string `dynamodbav:"deliveryLocation"`
	OrderStatus      string `dynamodbav:"orderStatus"`
	CreatedTimestamp string `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string `dynamodbav:"updatedTimestamp"`
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) String() string {
	return fmt.Sprintf("ID: {%s}, PaymentMethod: {%s}, Submitted: {%s}",
		o.PK,
		o.PaymentMethod,
		o.DeliveryLocation,
	)
}
