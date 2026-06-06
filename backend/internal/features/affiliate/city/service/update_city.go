package city_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *CityService) UpdateCity(
	ctx context.Context,
	id int,
	city core_transport_dto.UpdateCityRequest,
) (domain.City, error) {
	if err := city.Validate(); err != nil {
		return domain.City{}, fmt.Errorf("validate update city request: %w", err)
	}

	domainCity, err := s.cityRepository.UpdateCity(ctx, id, city)
	if err != nil {
		return domain.City{}, fmt.Errorf("update city: %w", err)
	}

	return domainCity, nil
}