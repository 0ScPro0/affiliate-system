package core_transport_dto

import "time"

type CreateOfferRequest struct {
	PartnerID   int       `json:"partner_id" validate:"required"`
	CategoryID  int       `json:"category_id" validate:"required"`
	CityID      int       `json:"city_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=1,max=100"`
	Description *string   `json:"description" validate:"omitempty,min=1,max=1000"`
	ExpireAt 	time.Time `json:"expire_at" validate:"required"`
}

type UpdateOfferRequest struct {
	ID          int        `json:"id" validate:"required"`
	Name        *string    `json:"name" validate:"omitempty,min=1,max=100"`
	Description *string    `json:"description" validate:"omitempty,min=1,max=1000"`
	ExpireAt    *time.Time `json:"expire_at" validate:"omitempty"`
}

type OfferResponse struct {
	ID          int       `json:"id"`
	PartnerID   int       `json:"partner_id"`
	CategoryID  int       `json:"category_id"`
	CityID      int       `json:"city_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ExpireAt 	time.Time `json:"expire_at"`
}