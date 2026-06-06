package core_transport_dto

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type CreateCityRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

func (r *CreateCityRequest) Validate() error {
	if !core_utils.ValidateStringLen(r.Name, 1, 50) {
		return fmt.Errorf(
			"invalid `Name` length: %d (must be 1-50): %w",
			core_utils.GetStringLen(r.Name),
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type UpdateCityRequest struct {
	Name *string `json:"name" validate:"omitempty,min=1,max=50"`
}

func (r *UpdateCityRequest) Validate() error {
	if r.Name != nil {
		if !core_utils.ValidateStringLen(*r.Name, 1, 50) {
			return fmt.Errorf(
				"invalid `Name` length: %d (must be 1-50): %w",
				core_utils.GetStringLen(*r.Name),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type CityResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}