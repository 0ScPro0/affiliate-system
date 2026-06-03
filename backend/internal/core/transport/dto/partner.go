package core_transport_dto

import "time"

type CreatePartnerRequest struct {
	Name        string
	Description *string
}

type UpdatePartnerRequest struct {
	ID          int
	Name        *string
	Description *string
}

type PartnerResponse struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}