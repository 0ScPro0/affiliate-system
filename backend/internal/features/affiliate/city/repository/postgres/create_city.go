package city_postgres_repository

import (
	"context"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *CityRepository) CreateCity(
	ctx context.Context,
	city core_transport_dto.CreateCityRequest,
) (domain.City, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		INSERT INTO affiliate_system.cities (name)
		VALUES ($1)
		RETURNING id, name, created_at;
	`

	row := r.pool.QueryRow(ctx, query, city.Name)

	var cityModel core_database_models.CityModel
	err := row.Scan(
		&cityModel.ID,
		&cityModel.Name,
		&cityModel.CreatedAt,
	)
	if err != nil {
		return domain.City{}, fmt.Errorf("create error: %w", err)
	}

	cityDomain := domain.NewCity(
		cityModel.ID,
		cityModel.Name,
		cityModel.CreatedAt,
	)

	return cityDomain, nil
}