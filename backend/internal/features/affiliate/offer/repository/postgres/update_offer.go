package offer_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *OfferRepository) UpdateOffer(
	ctx context.Context,
	id int,
	offer core_transport_dto.UpdateOfferRequest,
) (domain.Offer, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()
	
	query := `
		UPDATE affiliate_system.offers
		SET name = COALESCE(NULLIF($2, ''), name),
			description = description = COALESCE(NULLIF($3, ''), description),
			expire_at = COALESCE(NULLIF($4, ''), expire_at)
		WHERE id = $1
		RETURNING id, partner_id, category_id, city_id, name, description, created_at, expire_at
	`
	
	row := r.pool.QueryRow(ctx, query, id, offer.Name, offer.Description, offer.ExpireAt)

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
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Offer{}, fmt.Errorf("offer with id %d not found", id)
		}
		return domain.Offer{}, fmt.Errorf("update error: %w", err)
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