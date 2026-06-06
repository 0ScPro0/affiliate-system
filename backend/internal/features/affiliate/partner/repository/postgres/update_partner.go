package partner_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *PartnerRepository) UpdatePartner(
	ctx context.Context,
	id int,
	partner core_transport_dto.UpdatePartnerRequest,
) (domain.Partner, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()
	
	query := `
		UPDATE affiliate_system.partners
		SET name = COALESCE(NULLIF($2, ''), name),
			description = COALESCE(NULLIF($3, ''), description)
		END
		WHERE id = $1
		RETURNING id, name, description, created_at
	`
	
	row := r.pool.QueryRow(ctx, query, id, partner.Name, partner.Description)

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
		return domain.Partner{}, fmt.Errorf("update error: %w", err)
	}

	partnerDomain := domain.NewPartner(
		partnerModel.ID,
		partnerModel.Name,
		partnerModel.Description,
		partnerModel.CreatedAt,
	)

	return partnerDomain, nil
}