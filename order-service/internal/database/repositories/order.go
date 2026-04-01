package repositories

import (
	"assign1/internal/database/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}

func (r *OrderRepository) Create(ctx context.Context, order *models.Order) (*models.Order, error) {
	query := `
		INSERT INTO orders (customer_id, item_name, amount)
		VALUES ($1, $2, $3)
		RETURNING id, status, created_at
	`

	err := r.pool.QueryRow(ctx, query,
		order.CustomerID,
		order.ItemName,
		order.Amount,
	).Scan(&order.ID, &order.Status, &order.CreatedAt)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, id string) (*models.Order, error) {
	query := `
		SELECT id, customer_id, item_name, amount, status, created_at
		FROM orders
		WHERE id = $1
	`

	o := &models.Order{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&o.ID,
		&o.CustomerID,
		&o.ItemName,
		&o.Amount,
		&o.Status,
		&o.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, id string, status models.OrderStatus) error {
	query := `
		UPDATE orders
		SET status = $1
		WHERE id = $2
	`

	_, err := r.pool.Exec(ctx, query, status, id)
	return err
}
