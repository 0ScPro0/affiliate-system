package user_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		DELETE FROM affiliate_system.users 
		WHERE id = $1
	`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}