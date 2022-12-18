package models

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/spf13/viper"
)

type Order struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	Type             string `dynamodbav:"type"`
	CustomerID       uint32 `dynamodbav:"customerId"`
	SellerID         uint32 `dynamodbav:"sellerId"`
	PaymentMethod    string `dynamodbav:"paymentMethod"`
	DeliveryLocation string `dynamodbav:"deliveryLocation"`
	OrderStatus      string `dynamodbav:"orderStatus"`
	CreatedTimestamp string `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string `dynamodbav:"updatedTimestamp"`
}

type OrderRepo struct {
	db *dynamodb.Client
}

func NewOrderRepo(db *dynamodb.Client) *OrderRepo {
	return &OrderRepo{db}
}

func (o OrderRepo) CreateOrder(order Order) (*pb.OrderResponse, error) {
	av, err := attributevalue.MarshalMap(order)

	if err != nil {
		return nil, errors.New("error while marshelling")
	}

	_, err = o.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(viper.GetString("DB_TABLE")),
		Item:      av,
	})

	if err != nil {
		return nil, errors.New("error while creating data")
	}

	return &pb.OrderResponse{OrderId: order.PK}, nil
}
