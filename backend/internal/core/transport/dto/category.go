package core_transport_dto

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type CreateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
}

func (r *CreateCategoryRequest) Validate() error {
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

type UpdateCategoryRequest struct {
	ID          int     `json:"id" validate:"required"`
	Name        *string `json:"name" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
}

func (r *UpdateCategoryRequest) Validate() error {
	if r.ID <= 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", r.ID, core_errors.ErrInvalidArgument)
	}

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

type CategoryResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}