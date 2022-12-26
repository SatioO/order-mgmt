package service

import (
	"github.com/satioO/order-mgmt/config"
	commands "github.com/satioO/order-mgmt/internal/order/commands/v1"
	queries "github.com/satioO/order-mgmt/internal/order/queries/v1"
	"github.com/satioO/order-mgmt/pkg/es"
	"github.com/satioO/order-mgmt/pkg/logger"
)

type OrderService struct {
	Commands *commands.OrderCommands
	Queries  *queries.OrderQueries
}

func NewOrderService(log logger.Logger, cfg *config.Config, store es.AggregateStore) *OrderService {
	createOrderHandler := commands.NewCreateOrderHandler(log, cfg, store)

	commands := commands.NewOrderCommands(
		createOrderHandler,
	)

	return &OrderService{Commands: commands, Queries: nil}
}
