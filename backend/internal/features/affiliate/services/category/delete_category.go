package category_service

import (
	"context"
	"fmt"
)

func (s *CategoryService) DeleteCategory(
	ctx context.Context,
	id int,
) error {
	if err := s.categoryRepository.DeleteCategory(ctx, id); err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	return nil
}