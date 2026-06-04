package category_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (r *CategoryRepository) GetCategoryByID(
	ctx context.Context, 
	id int,
) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, created_at
		FROM affiliate_system.categories
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var categoryModel core_database_models.CategoryModel
	err := row.Scan(
		&categoryModel.ID,
		&categoryModel.Name,
		&categoryModel.Description,
		&categoryModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, fmt.Errorf("category with id %d not found", id)
		}
		return domain.Category{}, fmt.Errorf("get by id error: %w", err)
	}

	categoryDomain := domain.NewCategory(
		categoryModel.ID,
		categoryModel.Name,
		categoryModel.Description,
		categoryModel.CreatedAt,
	)

	return categoryDomain, nil
}