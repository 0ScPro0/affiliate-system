package category_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *CategoryService) UpdateCategory(
	ctx context.Context,
	category core_transport_dto.UpdateCategoryRequest,
) (domain.Category, error) {
	if err := category.Validate(); err != nil {
		return domain.Category{}, fmt.Errorf("validate update category request: %w", err)
	}

	domainCategory, err := s.categoryRepository.UpdateCategory(ctx, category)
	if err != nil {
		return domain.Category{}, fmt.Errorf("update category: %w", err)
	}

	return domainCategory, nil
}