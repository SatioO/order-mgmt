package models

type Product struct {
	ProductID   string `dynamodbav:"productID"`
	ProductName string `dynamodbav:"productName"`
}

type Order struct {
	PK               string  `dynamodbav:"PK"`
	SK               string  `dynamodbav:"SK"`
	CustomerID       string  `dynamodbav:"customerId"`
	SellerID         string  `dynamodbav:"sellerId"`
	Product          Product `dynamodbav:"product"`
	PaymentMethod    string  `dynamodbav:"paymentMethod"`
	DeliveryLocation string  `dynamodbav:"deliveryLocation"`
	OrderStatus      string  `dynamodbav:"orderStatus"`
	CreatedTimestamp string  `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string  `dynamodbav:"updatedTimestamp"`
}

func CreateOrder() (*Order, error) {
	return nil, nil
}
