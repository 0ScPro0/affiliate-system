package category_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *CategoryRepository) UpdateCategory(
	ctx context.Context,
	category core_transport_dto.UpdateCategoryRequest,
) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()
	
	query := `
		UPDATE affiliate_system.categories
		SET name = CASE 
			WHEN $2 IS NOT NULL AND $2 != '' THEN $2 
			ELSE name 
		END,
		description = CASE 
			WHEN $3 IS NOT NULL AND $3 != '' THEN $3 
			ELSE description 
		END
		WHERE id = $1
		RETURNING id, name, description, created_at
	`
	
	row := r.pool.QueryRow(ctx, query, category.ID, category.Name, category.Description)

	var categoryModel core_database_models.CategoryModel
	err := row.Scan(
		&categoryModel.ID,
		&categoryModel.Name,
		&categoryModel.Description,
		&categoryModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, fmt.Errorf("category with id %d not found", category.ID)
		}
		return domain.Category{}, fmt.Errorf("update error: %w", err)
	}

	categoryDomain := domain.NewCategory(
		categoryModel.ID,
		categoryModel.Name,
		categoryModel.Description,
		categoryModel.CreatedAt,
	)

	return categoryDomain, nil
}