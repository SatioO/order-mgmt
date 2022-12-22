package main

import (
	"log"
	"net"

	"github.com/satioO/order-mgmt/config"
	delivery "github.com/satioO/order-mgmt/internal/order/delivery/grpc"
	"github.com/satioO/order-mgmt/internal/order/service"
	"github.com/satioO/order-mgmt/proto"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	// dbCon := db.Init()
	// queueCon := queue.Init()

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()

	// order service registration
	// orderRepo := models.NewOrderRepo(dbCon, c.DynamoDBTable)
	// orderItemRepo := models.NewOrderItemRepo(dbCon, c.DynamoDBTable)
	// orderSvc := services.NewOrderService(queueCon, orderRepo, orderItemRepo)
	// proto.RegisterOrderServiceServer(grpcServer, orderSvc)
	orderSvc := service.NewOrderService(&c)
	orderGrpcSvc := delivery.NewOrderGRPCService(orderSvc)
	proto.RegisterOrderServiceServer(grpcServer, orderGrpcSvc)

	log.Printf("Order Mgmt Service is running at PORT %s", c.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
