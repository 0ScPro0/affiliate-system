package offer_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *OfferService) CreateOffer(
	ctx context.Context,
	offer core_transport_dto.CreateOfferRequest,
) (domain.Offer, error) {
	if err := offer.Validate(); err != nil {
		return domain.Offer{}, fmt.Errorf("validate create offer request: %w", err)
	}

	domainOffer, err := s.offerRepository.CreateOffer(ctx, offer)
	if err != nil {
		return domain.Offer{}, fmt.Errorf("create offer: %w", err)
	}

	return domainOffer, nil
}