package models

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type OrderItem struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	GSI1PK           string `dynamodbav:"GSI1-PK"`
	GSI1SK           string `dynamodbav:"GSI1-SK"`
	Type             string `dynamodbav:"type"`
	ProductName      string `dynamodbav:"productName"`
	CreatedTimestamp string `dynamodbav:"createdTimestamp"`
	UpdatedTimestamp string `dynamodbav:"updatedTimestamp"`
}

type OrderItemRepo struct {
	db        *dynamodb.Client
	tableName string
}

func NewOrderItemRepo(db *dynamodb.Client, tableName string) *OrderItemRepo {
	return &OrderItemRepo{db, tableName}
}

func (o OrderItemRepo) CreateOrderItem(orderItem OrderItem) error {
	av, err := attributevalue.MarshalMap(orderItem)

	if err != nil {
		return err
	}

	_, err = o.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(o.tableName),
		Item:      av,
	})

	return err
}

func (o OrderItemRepo) BatchWriteOrderItems(orderItems []OrderItem) error {
	var writeItems []types.WriteRequest

	for _, orderItem := range orderItems {
		av, err := attributevalue.MarshalMap(orderItem)
		if err != nil {
			return err
		}

		writeItems = append(writeItems, types.WriteRequest{PutRequest: &types.PutRequest{Item: av}})
	}

	_, err := o.db.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{o.tableName: writeItems},
	})

	return err
}
