package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	Order *OrderRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		Order: NewOrderRepository(pool),
	}
}
