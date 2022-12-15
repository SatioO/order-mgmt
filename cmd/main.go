package main

import (
	"log"
	"net"

	"github.com/satioO/order-mgmt/pkg/config"
	"github.com/satioO/order-mgmt/pkg/db"
	"github.com/satioO/order-mgmt/pkg/pb"
	"github.com/satioO/order-mgmt/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	dbCon := db.Init()

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()

	// order service registration
	orderSvc := services.NewOrderService(dbCon)
	pb.RegisterOrderServiceServer(grpcServer, orderSvc)

	log.Printf("Order Mgmt Service is running at PORT %s", c.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
