package partner_postgres_repository

import (
	"context"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *PartnerRepository) CreatePartner(
	ctx context.Context,
	partner core_transport_dto.CreatePartnerRequest,
) (domain.Partner, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		INSERT INTO affiliate_system.partners (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at
	`

	row := r.pool.QueryRow(ctx, query, partner.Name, partner.Description)

	var partnerModel core_database_models.PartnerModel
	err := row.Scan(
		&partnerModel.ID,
		&partnerModel.Name,
		&partnerModel.Description,
		&partnerModel.CreatedAt,
	)
	if err != nil {
		return domain.Partner{}, fmt.Errorf("create error: %w", err)
	}

	partnerDomain := domain.NewPartner(
		partnerModel.ID,
		partnerModel.Name,
		partnerModel.Description,
		partnerModel.CreatedAt,
	)

	return partnerDomain, nil
}