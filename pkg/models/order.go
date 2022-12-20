package models

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/satioO/order-mgmt/pkg/pb"
)

type Order struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	GSI1PK           string `dynamodbav:"GSI1-PK"`
	GSI1SK           string `dynamodbav:"GSI1-SK"`
	Type             string `dynamodbav:"type"`
	CustomerID       uint32 `dynamodbav:"customerId"`
	SellerID         uint32 `dynamodbav:"sellerId"`
	OrderID          string `dynamodbav:"orderId"`
	PaymentMethod    string `dynamodbav:"paymentMethod"`
	DeliveryLocation string `dynamodbav:"deliveryLocation"`
	OrderStatus      string `dynamodbav:"orderStatus"`
	CreatedTimestamp string `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string `dynamodbav:"updatedTimestamp"`
}

type OrderRepo struct {
	db        *dynamodb.Client
	tableName string
}

func NewOrderRepo(db *dynamodb.Client, tableName string) *OrderRepo {
	return &OrderRepo{db, tableName}
}

func (o OrderRepo) GetOrders(ctx context.Context) (*pb.GetOrdersResponse, error) {
	result, err := o.db.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(o.tableName),
		IndexName:              aws.String("GSI1"),
		KeyConditionExpression: aws.String("#key = :key"),
		ExpressionAttributeNames: map[string]string{
			"#key": "GSI1-PK",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":key": &types.AttributeValueMemberS{Value: "ORDER"},
		},
	})

	if err != nil {
		return nil, err
	}

	orders := pb.GetOrdersResponse{}
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &orders.Orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (o OrderRepo) CreateOrder(ctx context.Context, order Order) (*pb.OrderResponse, error) {
	av, err := attributevalue.MarshalMap(order)

	if err != nil {
		return nil, errors.New("error while marshelling")
	}

	_, err = o.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(o.tableName),
		Item:      av,
	})

	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{OrderId: order.OrderID}, nil
}
