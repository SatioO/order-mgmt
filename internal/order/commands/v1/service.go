package v1

type OrderCommands struct {
	CreateOrder CreateOrderCommandHandler
}

func NewOrderCommands(createOrder CreateOrderCommandHandler) *OrderCommands {
	return &OrderCommands{
		CreateOrder: createOrder,
	}
}
