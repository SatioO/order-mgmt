package models

type Order struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	GSI1PK           string `dynamodbav:"GSI1-PK"`
	GSI1SK           string `dynamodbav:"GSI1-SK"`
	Type             string `dynamodbav:"type"`
	OrderID          string `dynamodbav:"orderId"`
	CustomerID       string `dynamodbav:"customerId"`
	PaymentMethod    string `dynamodbav:"paymentMethod"`
	DeliveryLocation string `dynamodbav:"deliveryLocation"`
	OrderStatus      string `dynamodbav:"orderStatus"`
	CreatedTimestamp string `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string `dynamodbav:"updatedTimestamp"`
}
