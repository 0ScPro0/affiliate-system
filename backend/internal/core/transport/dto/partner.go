package core_transport_dto

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type CreatePartnerRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
}

func (r *CreatePartnerRequest) Validate() error {
	if !core_utils.ValidateStringLen(r.Name, 1, 100) {
		return fmt.Errorf(
			"invalid `Name` length: %d (must be 1-100): %w",
			core_utils.GetStringLen(r.Name),
			core_errors.ErrInvalidArgument,
		)
	}

	if r.Description != nil {
		if !core_utils.ValidateStringLen(*r.Description, 1, 1000) {
			return fmt.Errorf(
				"invalid `Description` length: %d (must be 1-1000): %w",
				core_utils.GetStringLen(*r.Description),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type UpdatePartnerRequest struct {
	Name        *string `json:"name" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
}

func (r *UpdatePartnerRequest) Validate() error {
	if r.Name != nil {
		if !core_utils.ValidateStringLen(*r.Name, 1, 100) {
			return fmt.Errorf(
				"invalid `Name` length: %d (must be 1-100): %w",
				core_utils.GetStringLen(*r.Name),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Description != nil {
		if !core_utils.ValidateStringLen(*r.Description, 1, 1000) {
			return fmt.Errorf(
				"invalid `Description` length: %d (must be 1-1000): %w",
				core_utils.GetStringLen(*r.Description),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type PartnerResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}