package offer_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (s *OfferService) GetOfferByID(
	ctx context.Context,
	id int,
) (domain.Offer, error) {
	domainOffer, err := s.offerRepository.GetOfferByID(ctx, id)
	if err != nil {
		return domain.Offer{}, fmt.Errorf("get offer by id: %w", err)
	}

	return domainOffer, nil
}