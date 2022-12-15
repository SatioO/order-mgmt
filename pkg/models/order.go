package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/spf13/viper"
)

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

type OrderRepo struct {
	db *dynamodb.Client
}

func NewOrderRepo(db *dynamodb.Client) *OrderRepo {
	return &OrderRepo{db}
}

func (o OrderRepo) CreateOrder(body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	orderId, _ := uuid.NewUUID()

	_, err := o.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(viper.GetString("DB_TABLE")),
		Item: map[string]types.AttributeValue{
			"PK":               &types.AttributeValueMemberS{Value: fmt.Sprintf("ORDER#%s", orderId)},
			"SK":               &types.AttributeValueMemberS{Value: fmt.Sprintf("ORDER#%s", orderId)},
			"customerId":       &types.AttributeValueMemberS{Value: body.DeliveryLocation},
			"sellerId":         &types.AttributeValueMemberN{Value: strconv.Itoa(int(body.SellerId))},
			"paymentMethod":    &types.AttributeValueMemberS{Value: body.PaymentMethod.Enum().String()},
			"deliveryLocation": &types.AttributeValueMemberS{Value: body.DeliveryLocation},
			"orderStatus":      &types.AttributeValueMemberS{Value: "ORDER_CREATED"},
			"createdTimestamp": &types.AttributeValueMemberS{Value: time.Now().String()},
			"updatedTimestamp": &types.AttributeValueMemberS{Value: time.Now().String()},
		},
	})

	return &pb.OrderResponse{OrderId: orderId.String()}, err
}
