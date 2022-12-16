package models

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/viper"
)

type OrderItem struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	ProductName      string `dynamodbav:"productName"`
	CreatedTimestamp string `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string `dynamodbav:"updatedTimestamp"`
}

type OrderItemRepo struct {
	db *dynamodb.Client
}

func NewOrderItemRepo(db *dynamodb.Client) *OrderItemRepo {
	return &OrderItemRepo{db}
}

func (o OrderItemRepo) CreateOrderItem(orderItem OrderItem) error {
	av, err := attributevalue.MarshalMap(orderItem)

	if err != nil {
		return err
	}

	_, err = o.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(viper.GetString("DB_TABLE")),
		Item:      av,
	})

	return err
}
