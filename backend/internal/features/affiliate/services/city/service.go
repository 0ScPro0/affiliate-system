package city_service

import (
	"context"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

type CityService struct {
	cityRepository CityRepository
}

type CityRepository interface {
	CreateCity(
		ctx context.Context,
		city core_transport_dto.CreateCityRequest,
	) (domain.City, error)

	GetCityByID(
		ctx context.Context, 
		id int,
	) (domain.City, error)

	UpdateCity(
		ctx context.Context,
		city core_transport_dto.UpdateCityRequest,
	) (domain.City, error)

	DeleteCity(
		ctx context.Context, 
		id int,
	) error
}

func NewCityService(
	cityRepository CityRepository, 
) *CityService {
	return &CityService{
		cityRepository: cityRepository,
	}
}