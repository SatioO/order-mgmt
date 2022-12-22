package main

import (
	"log"
	"net"

	"github.com/satioO/order-mgmt/config"
	delivery "github.com/satioO/order-mgmt/internal/order/delivery/grpc"
	"github.com/satioO/order-mgmt/internal/order/service"
	"github.com/satioO/order-mgmt/pkg/logger"
	"github.com/satioO/order-mgmt/proto"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	appLogger := logger.NewAppLogger(c)
	appLogger.InitLogger()
	appLogger.WithName(c.ServerName)

	// dbCon := db.Init()
	// queueCon := queue.Init()

	lis, err := net.Listen("tcp", c.GRPC.Port)
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()

	// order service registration
	orderSvc := service.NewOrderService(c)
	orderGrpcSvc := delivery.NewOrderGRPCService(appLogger, orderSvc)
	proto.RegisterOrderServiceServer(grpcServer, orderGrpcSvc)

	log.Printf("%s gRPC server is listening on port: {%s}", c.ServerName, c.GRPC.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
