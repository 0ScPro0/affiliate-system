package category_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (s *CategoryService) GetCategoryByID(
	ctx context.Context,
	id int,
) (domain.Category, error) {
	domainCategory, err := s.categoryRepository.GetCategoryByID(ctx, id) 
	if err != nil {
		return domain.Category{}, fmt.Errorf("get category by id: %w", err)
	}
	
	return domainCategory, nil
}