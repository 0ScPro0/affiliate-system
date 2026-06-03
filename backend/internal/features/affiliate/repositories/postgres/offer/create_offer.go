package offer_postgres_repository

import (
	"context"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *OfferRepository) CreateOffer(
	ctx context.Context,
	offer core_transport_dto.CreateOfferRequest,
) (domain.Offer, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		INSERT INTO affiliate_system.offers (partner_id, category_id, city_id, name, description, expire_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, partner_id, category_id, city_id, name, description, created_at, expire_at
	`

	row := r.pool.QueryRow(
		ctx, query,
		offer.PartnerID,
		offer.CategoryID,
		offer.CityID,
		offer.Name,
		offer.Description,
		offer.ExpireAt,
	)

	var offerModel core_database_models.OfferModel
	err := row.Scan(
		&offerModel.ID,
		&offerModel.PartnerID,
		&offerModel.CategoryID,
		&offerModel.CityID,
		&offerModel.Name,
		&offerModel.Description,
		&offerModel.CreatedAt,
		&offerModel.ExpireAt,
	)
	if err != nil {
		return domain.Offer{}, fmt.Errorf("create error: %w", err)
	}

	offerDomain := domain.NewOffer(
		offerModel.ID,
		offerModel.PartnerID,
		offerModel.CategoryID,
		offerModel.CityID,
		offerModel.Name,
		offerModel.Description,
		offerModel.CreatedAt,
		offerModel.ExpireAt,
	)

	return offerDomain, nil
}