package partner_postgres_repository

import core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"

type PartnerRepository struct {
	pool core_postgres_pool.Pool
}

func NewPartnerRepository(
	pool core_postgres_pool.Pool,
) *PartnerRepository {
	return &PartnerRepository{
		pool: pool,
	}
}