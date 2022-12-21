package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/satioO/order-mgmt/internal/mapper"
	"github.com/satioO/order-mgmt/internal/models"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/spf13/viper"
)

type orderService struct {
	pb.UnimplementedOrderServiceServer
	queue         *sqs.Client
	orderRepo     *models.OrderRepo
	orderItemRepo *models.OrderItemRepo
}

func NewOrderService(queue *sqs.Client, orderRepo *models.OrderRepo, orderItemRepo *models.OrderItemRepo) *orderService {
	return &orderService{
		queue:         queue,
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (o orderService) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	return o.orderRepo.GetOrders(ctx, req)
}

func (o orderService) CreateOrder(ctx context.Context, body *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	orderEntity := mapper.ToCreateOrderEntity(body)
	createdOrder, err := o.orderRepo.CreateOrder(ctx, orderEntity)

	if err != nil {
		return nil, err
	}

	var orderItems []models.OrderItem
	for _, product := range body.Products {
		orderItemEntity := mapper.ToCreateOrderItemEntity(orderEntity.SK, product)
		orderItems = append(orderItems, orderItemEntity)
	}

	if err := o.orderItemRepo.BatchWriteOrderItems(orderItems); err != nil {
		return nil, err
	}

	_, err = o.queue.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:     aws.String(viper.GetString("QUEUE_CREATE_ORDER_TOPIC")),
		MessageBody:  aws.String("ORDER:" + createdOrder.OrderId),
		DelaySeconds: *aws.Int32(0),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"orderId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(createdOrder.OrderId),
			},
			"customerId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(body.CustomerId),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return createdOrder, err
}
