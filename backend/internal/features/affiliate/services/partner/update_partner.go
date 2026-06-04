package partner_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *PartnerService) UpdatePartner(
	ctx context.Context,
	partner core_transport_dto.UpdatePartnerRequest,
) (domain.Partner, error) {
	if err := partner.Validate(); err != nil {
		return domain.Partner{}, fmt.Errorf("validate update partner request: %w", err)
	}

	domainPartner, err := s.partnerRepository.UpdatePartner(ctx, partner)
	if err != nil {
		return domain.Partner{}, fmt.Errorf("update partner: %w", err)
	}

	return domainPartner, nil
}