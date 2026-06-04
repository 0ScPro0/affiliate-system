package offer_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (r *OfferRepository) GetOfferByID(
	ctx context.Context, 
	id int,
) (domain.Offer, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		SELECT id, partner_id, category_id, city_id, name, description, created_at, expire_at
		FROM affiliate_system.offers
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

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
		return domain.Offer{}, fmt.Errorf("get by id error: %w", err)
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