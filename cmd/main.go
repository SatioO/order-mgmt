package main

import (
	"log"
	"net"

	"github.com/satioO/order-mgmt/pkg/config"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/satioO/order-mgmt/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()

	orderSvc := services.OrderService{}
	pb.RegisterOrderServiceServer(grpcServer, &orderSvc)

	log.Println("Order Mgmt Service is running at PORT 3000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
