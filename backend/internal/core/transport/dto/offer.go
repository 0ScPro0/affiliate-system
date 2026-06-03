package core_transport_dto

import "time"

type CreateOfferRequest struct {
	PartnerID   int
	CategoryID  int
	CityID      int
	Name        string
	Description *string
	ExpireAt 	time.Time
}

type UpdateOfferRequest struct {
	ID          int
	Name        *string
	Description *string
	ExpireAt    *time.Time
}

type OfferResponse struct {
	ID          int
	PartnerID   int
	CategoryID  int
	CityID      int
	Name        string
	Description *string
	CreatedAt   time.Time
	ExpireAt 	time.Time
}