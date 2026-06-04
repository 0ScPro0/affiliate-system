package partner_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *PartnerService) CreatePartner(
	ctx context.Context,
	partner core_transport_dto.CreatePartnerRequest,
) (domain.Partner, error) {
	if err := partner.Validate(); err != nil {
		return domain.Partner{}, fmt.Errorf("validate create partner request: %w", err)
	}

	domainPartner, err := s.partnerRepository.CreatePartner(ctx, partner)
	if err != nil {
		return domain.Partner{}, fmt.Errorf("create partner: %w", err)
	}

	return domainPartner, nil
}