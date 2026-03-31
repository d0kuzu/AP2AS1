package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	Payment *PaymentRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		Payment: NewPaymentRepository(pool),
	}
}
