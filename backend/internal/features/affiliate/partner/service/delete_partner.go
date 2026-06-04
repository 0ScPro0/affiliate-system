package partner_service

import (
	"context"
	"fmt"
)

func (s *PartnerService) DeletePartner(
	ctx context.Context,
	id int,
) error {
	if err := s.partnerRepository.DeletePartner(ctx, id); err != nil {
		return fmt.Errorf("delete partner: %w", err)
	}

	return nil
}