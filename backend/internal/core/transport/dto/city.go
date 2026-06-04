package core_transport_dto

import "time"

type CreateCityRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

type UpdateCityRequest struct {
	ID        int     `json:"id" validate:"required"`
	Name      *string `json:"name" validate:"omitempty,min=1,max=50"`
}

type CityResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}