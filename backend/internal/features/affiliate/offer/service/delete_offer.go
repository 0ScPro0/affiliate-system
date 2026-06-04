package offer_service

import (
	"context"
	"fmt"
)

func (s *OfferService) DeleteOffer(
	ctx context.Context,
	id int,
) error {
	if err := s.offerRepository.DeleteOffer(ctx, id); err != nil {
		return fmt.Errorf("delete offer: %w", err)
	}

	return nil
}