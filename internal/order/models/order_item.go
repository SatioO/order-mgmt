package models

type OrderItem struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Quantity uint32  `json:"quantity"`
	Price    float32 `json:"price"`
}
