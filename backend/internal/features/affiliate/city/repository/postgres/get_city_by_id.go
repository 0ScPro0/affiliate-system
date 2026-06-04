package city_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (r *CityRepository) GetCityByID(
	ctx context.Context, 
	id int,
) (domain.City, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		SELECT id, name, created_at
		FROM affiliate_system.cities
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var cityModel core_database_models.CityModel
	err := row.Scan(
		&cityModel.ID,
		&cityModel.Name,
		&cityModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.City{}, fmt.Errorf("city with id %d not found", id)
		}
		return domain.City{}, fmt.Errorf("get by id error: %w", err)
	}

	cityDomain := domain.NewCity(
		cityModel.ID,
		cityModel.Name,
		cityModel.CreatedAt,
	)

	return cityDomain, nil
}