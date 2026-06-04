package user_postgres_repository

import core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"

type UserRepository struct {
	pool core_postgres_pool.Pool
}

func NewUserRepository(
	pool core_postgres_pool.Pool,
) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}