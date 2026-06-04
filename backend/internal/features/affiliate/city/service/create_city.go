package city_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *CityService) CreateCity(
	ctx context.Context,
	city core_transport_dto.CreateCityRequest,
) (domain.City, error) {
	if err := city.Validate(); err != nil {
		return domain.City{}, fmt.Errorf("validate city create request: %w", err)
	}

	cityDomain, err := s.cityRepository.CreateCity(ctx, city)
	if err != nil {
		return domain.City{}, fmt.Errorf("create city: %w", err)
	}

	return cityDomain, nil
}