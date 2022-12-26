package v1

import (
	"context"
	"log"

	"github.com/satioO/order-mgmt/config"
	"github.com/satioO/order-mgmt/internal/order/aggregate"
	"github.com/satioO/order-mgmt/pkg/es"
	"github.com/satioO/order-mgmt/pkg/logger"
)

type CreateOrderCommandHandler interface {
	Handle(ctx context.Context, command *CreateOrderCommand) error
}

type createOrderHandler struct {
	log   logger.Logger
	cfg   *config.Config
	store es.AggregateStore
}

func NewCreateOrderHandler(log logger.Logger, cfg *config.Config, store es.AggregateStore) *createOrderHandler {
	return &createOrderHandler{log, cfg, store}
}

func (c *createOrderHandler) Handle(ctx context.Context, command *CreateOrderCommand) error {
	c.log.Infof("AggregateID: %s", command.GetAggregateID())
	order := aggregate.NewOrderAggregateWithID(command.AggregateID)
	
	log.Println(order.ID, order.Version)

	return nil
}
