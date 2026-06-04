package offer_postgres_repository

import core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"

type OfferRepository struct {
	pool core_postgres_pool.Pool
}

func NewOfferRepository(
	pool core_postgres_pool.Pool,
) *OfferRepository {
	return &OfferRepository{
		pool: pool,
	}
}