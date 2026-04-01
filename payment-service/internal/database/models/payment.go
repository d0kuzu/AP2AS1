package models

import "time"

type PaymentStatus string

const (
	PaymentAuthorized PaymentStatus = "Authorized"
	PaymentDeclined   PaymentStatus = "Declined"
)

type Payment struct {
	ID            string        `json:"id"`
	OrderID       string        `json:"order_id"`
	TransactionID string        `json:"transaction_id"`
	Amount        int64         `json:"amount"`
	Status        PaymentStatus `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
}
