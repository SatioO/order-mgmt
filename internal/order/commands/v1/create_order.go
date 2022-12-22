package v1

import (
	"context"
	"log"

	"github.com/satioO/order-mgmt/config"
)

type CreateOrderCommandHandler interface {
	Handle(ctx context.Context, command *CreateOrderCommand) error
}

type createOrderHandler struct {
	cfg *config.Config
}

func NewCreateOrderHandler(cfg *config.Config) *createOrderHandler {
	return &createOrderHandler{cfg}
}

func (c *createOrderHandler) Handle(ctx context.Context, command *CreateOrderCommand) error {
	log.Println(command)
	return nil
}
