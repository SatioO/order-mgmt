package models

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/spf13/viper"
)

type Product struct {
	ProductID   int32  `dynamodbav:"productId"`
	ProductName string `dynamodbav:"productName"`
}

type Order struct {
	PK               string  `dynamodbav:"PK"`
	SK               string  `dynamodbav:"SK"`
	CustomerID       int32   `dynamodbav:"customerId"`
	SellerID         int32   `dynamodbav:"sellerId"`
	Product          Product `dynamodbav:"product"`
	PaymentMethod    string  `dynamodbav:"paymentMethod"`
	DeliveryLocation string  `dynamodbav:"deliveryLocation"`
	OrderStatus      string  `dynamodbav:"orderStatus"`
	CreatedTimestamp string  `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string  `dynamodbav:"updatedTimestamp"`
}

type OrderRepo struct {
	db *dynamodb.Client
}

func NewOrderRepo(db *dynamodb.Client) *OrderRepo {
	return &OrderRepo{db}
}

func (o OrderRepo) CreateOrder(body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	orderId, _ := uuid.NewUUID()

	order := Order{
		PK:         fmt.Sprintf("ORDER#%s", orderId),
		SK:         fmt.Sprintf("ORDER#%s", orderId),
		CustomerID: body.CustomerId,
		SellerID:   body.SellerId,
		Product: Product{
			ProductID:   body.Product.ProductId,
			ProductName: body.Product.ProductName,
		},
		PaymentMethod:    body.PaymentMethod.String(),
		DeliveryLocation: body.DeliveryLocation,
		OrderStatus:      "ORDER_CREATED",
		CreatedTimestamp: time.Now().String(),
		UpdatedTimestamp: time.Now().String(),
	}

	av, err := attributevalue.MarshalMap(order)

	if err != nil {
		return nil, err
	}

	_, err = o.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(viper.GetString("DB_TABLE")),
		Item:      av,
	})

	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{OrderId: orderId.String()}, nil
}
