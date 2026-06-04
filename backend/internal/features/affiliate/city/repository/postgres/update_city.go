package city_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *CityRepository) UpdateCity(
	ctx context.Context,
	city core_transport_dto.UpdateCityRequest,
) (domain.City, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()
	
	query := `
		UPDATE affiliate_system.cities
		SET name = CASE 
			WHEN $2 IS NOT NULL AND $2 != '' THEN $2 
			ELSE name 
		END
		WHERE id = $1
		RETURNING id, name, created_at
	`
	
	row := r.pool.QueryRow(ctx, query, city.ID, city.Name)

	var cityModel core_database_models.CityModel
	err := row.Scan(
		&cityModel.ID,
		&cityModel.Name,
		&cityModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.City{}, fmt.Errorf("city with id %d not found", city.ID)
		}
		return domain.City{}, fmt.Errorf("update error: %w", err)
	}

	domainCity := domain.NewCity(
		cityModel.ID,
		cityModel.Name,
		cityModel.CreatedAt,
	)

	return domainCity, nil
}