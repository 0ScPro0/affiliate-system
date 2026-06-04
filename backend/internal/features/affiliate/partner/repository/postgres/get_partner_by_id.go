package partner_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (r *PartnerRepository) GetPartnerByID(
	ctx context.Context, 
	id int,
) (domain.Partner, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, created_at
		FROM affiliate_system.partners
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var partnerModel core_database_models.PartnerModel
	err := row.Scan(
		&partnerModel.ID,
		&partnerModel.Name,
		&partnerModel.Description,
		&partnerModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Partner{}, fmt.Errorf("partner with id %d not found", id)
		}
		return domain.Partner{}, fmt.Errorf("get by id error: %w", err)
	}

	partnerDomain := domain.NewPartner(
		partnerModel.ID,
		partnerModel.Name,
		partnerModel.Description,
		partnerModel.CreatedAt,
	)

	return partnerDomain, nil
}