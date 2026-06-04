package city_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (s *CityService) GetCityByID(
	ctx context.Context,
	id int,
) (domain.City, error) {
	domainCity, err := s.cityRepository.GetCityByID(ctx, id)
	if err != nil {
		return domain.City{}, fmt.Errorf("get city by id: %w", err)
	}

	return domainCity, nil
}