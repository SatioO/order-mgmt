package service

import (
	"github.com/satioO/order-mgmt/config"
	commands "github.com/satioO/order-mgmt/internal/order/commands/v1"
	queries "github.com/satioO/order-mgmt/internal/order/queries/v1"
)

type OrderService struct {
	Commands *commands.OrderCommands
	Queries  *queries.OrderQueries
}

func NewOrderService(cfg *config.Config) *OrderService {
	createOrderHandler := commands.NewCreateOrderHandler(cfg)

	commands := commands.NewOrderCommands(
		createOrderHandler,
	)

	return &OrderService{Commands: commands, Queries: nil}
}
