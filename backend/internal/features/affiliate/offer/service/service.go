package offer_service

import (
	"context"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

type OfferService struct {
	offerRepository OfferRepository
}

type OfferRepository interface {
	CreateOffer(
		ctx context.Context,
		offer core_transport_dto.CreateOfferRequest,
	) (domain.Offer, error)

	GetOfferByID(
		ctx context.Context,
		id int,
	) (domain.Offer, error)

	UpdateOffer(
		ctx context.Context,
		id int,
		ffer core_transport_dto.UpdateOfferRequest,
	) (domain.Offer, error)

	DeleteOffer(
		ctx context.Context,
		id int,
	) error
}

func NewOfferService(
	offerRepository OfferRepository,
) *OfferService {
	return &OfferService{
		offerRepository: offerRepository,
	}
}