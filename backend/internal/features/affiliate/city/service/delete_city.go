package city_service

import (
	"context"
	"fmt"
)

func (s *CityService) DeleteCity(
	ctx context.Context,
	id int,
) error {
	if err := s.cityRepository.DeleteCity(ctx, id); err != nil {
		return fmt.Errorf("delete city: %w", err)
	}

	return nil
}