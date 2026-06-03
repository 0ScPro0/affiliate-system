package category_postgres_repository

import (
	"context"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *CategoryRepository) CreateCategory(
	ctx context.Context,
	category core_transport_dto.CreateCategoryRequest,
) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		INSERT INTO affiliate_system.categories (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at;
	`

	row := r.pool.QueryRow(ctx, query, category.Name, category.Description)

	var categoryModel core_database_models.CategoryModel
	err := row.Scan(
		&categoryModel.ID,
		&categoryModel.Name,
		&categoryModel.Description,
		&categoryModel.CreatedAt,
	)
	if err != nil {
		return domain.Category{}, fmt.Errorf("create error: %w", err)
	}

	categoryDomain := domain.NewCategory(
		categoryModel.ID,
		categoryModel.Name,
		categoryModel.Description,
		categoryModel.CreatedAt,
	)

	return categoryDomain, nil
}