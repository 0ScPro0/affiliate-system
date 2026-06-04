package category_service

import (
	"context"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

type CategoryService struct {
	categoryRepository CategoryRepository
}

type CategoryRepository interface {
	CreateCategory(
		ctx context.Context,
		category core_transport_dto.CreateCategoryRequest,
	) (domain.Category, error)

	GetCategoryByID(
		ctx context.Context,
		id int,
	) (domain.Category, error)

	UpdateCategory(
		ctx context.Context,
		category core_transport_dto.UpdateCategoryRequest,
	) (domain.Category, error)

	DeleteCategory(
		ctx context.Context,
		id int,
	) error
}

func NewCategoryService(
	categoryRepository CategoryRepository,
) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}