package repositories

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	ItemName   string    `json:"item_name"`
	Amount     int64     `json:"amount"` // Amount in cents (e.g., 1000 = $10.00)
	Status     string    `json:"status"` // "Pending", "Paid", "Failed", "Cancelled"
	CreatedAt  time.Time `json:"created_at"`
}

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}
