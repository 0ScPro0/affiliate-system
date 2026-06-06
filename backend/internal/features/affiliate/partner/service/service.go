package partner_service

import (
	"context"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

type PartnerService struct {
	partnerRepository PartnerRepository
}

type PartnerRepository interface {
	CreatePartner(
		ctx context.Context,
		partner core_transport_dto.CreatePartnerRequest,
	) (domain.Partner, error)

	GetPartnerByID(
		ctx context.Context,
		id int,
	) (domain.Partner, error)

	UpdatePartner(
		ctx context.Context,
		id int,
		partner core_transport_dto.UpdatePartnerRequest,
	) (domain.Partner, error)

	DeletePartner(
		ctx context.Context,
		id int,
	) error
}

func NewPartnerService(
	partnerRepository PartnerRepository,
) *PartnerService {
	return &PartnerService{
		partnerRepository: partnerRepository,
	}
}