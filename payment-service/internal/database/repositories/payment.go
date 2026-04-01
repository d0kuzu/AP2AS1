package repositories

import (
	"assign1/internal/database/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	pool *pgxpool.Pool
}

func NewPaymentRepository(pool *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{pool: pool}
}

func (r *PaymentRepository) Create(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	query := `
		INSERT INTO payments (order_id, transaction_id, amount, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	err := r.pool.QueryRow(ctx, query,
		payment.OrderID,
		payment.TransactionID,
		payment.Amount,
		payment.Status,
	).Scan(&payment.ID, &payment.CreatedAt)

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) GetByOrderID(ctx context.Context, orderID string) (*models.Payment, error) {
	query := `
		SELECT id, order_id, transaction_id, amount, status, created_at
		FROM payments
		WHERE order_id = $1
		LIMIT 1
	`

	p := &models.Payment{}
	err := r.pool.QueryRow(ctx, query, orderID).Scan(
		&p.ID,
		&p.OrderID,
		&p.TransactionID,
		&p.Amount,
		&p.Status,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}
