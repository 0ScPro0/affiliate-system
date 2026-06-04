package category_postgres_repository

import core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"

type CategoryRepository struct {
	pool core_postgres_pool.Pool
}

func NewCategoryRepository(
	pool core_postgres_pool.Pool,
) *CategoryRepository {
	return &CategoryRepository{
		pool: pool,
	}
}