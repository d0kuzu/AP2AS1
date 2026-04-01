package order

type CreateOrderRequest struct {
	CustomerID string `json:"customer_id"`
	ItemName   string `json:"item_name"`
	Amount     int64  `json:"amount"`
}

type PaymentRequest struct {
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

type PaymentResponse struct {
	Status string `json:"status"`
}
