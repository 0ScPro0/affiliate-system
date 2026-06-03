package category_postgres_repository

import (
	"context"
	"fmt"
)

func (r *CategoryRepository) DeleteCategory(
	ctx context.Context, 
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		DELETE FROM affiliate_system.categories 
		WHERE id = $1
	`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("category with id %d not found", id)
	}

	return nil
}