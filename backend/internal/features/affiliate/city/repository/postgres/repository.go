package city_postgres_repository

import core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"

type CityRepository struct {
	pool core_postgres_pool.Pool
}

func NewCityRepository(
	pool core_postgres_pool.Pool,
) *CityRepository {
	return &CityRepository{
		pool: pool,
	}
}