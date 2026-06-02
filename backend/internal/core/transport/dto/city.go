package core_transport_dto

import "time"

type CreateCityRequest struct {
	Name string
}

type UpdateCityRequest struct {
	ID        int
	Name      *string
}

type CityResponse struct {
	ID        int
	Name      string
	CreatedAt time.Time
}