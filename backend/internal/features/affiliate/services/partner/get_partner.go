package partner_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (s *PartnerService) GetPartnerByID(
	ctx context.Context,
	id int,
) (domain.Partner, error) {
	domainPartner, err := s.partnerRepository.GetPartnerByID(ctx, id)
	if err != nil {
		return domain.Partner{}, fmt.Errorf("get partner by id: %w", err)
	}

	return domainPartner, nil
}