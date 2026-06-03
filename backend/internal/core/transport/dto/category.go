package core_transport_dto

import "time"

type CreateCategoryRequest struct {
	Name        string
	Description *string
}

type UpdateCategoryRequest struct {
	ID          int
	Name        *string
	Description *string
}

type CategoryResponse struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}