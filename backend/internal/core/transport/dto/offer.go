package core_transport_dto

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type CreateOfferRequest struct {
	PartnerID   int       `json:"partner_id" validate:"required"`
	CategoryID  int       `json:"category_id" validate:"required"`
	CityID      int       `json:"city_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=1,max=100"`
	Description *string   `json:"description" validate:"omitempty,min=1,max=1000"`
	ExpireAt    time.Time `json:"expire_at" validate:"required"`
}

func (r *CreateOfferRequest) Validate() error {
	if r.PartnerID <= 0 {
		return fmt.Errorf("invalid `PartnerID`: %d: %w", r.PartnerID, core_errors.ErrInvalidArgument)
	}

	if r.CategoryID <= 0 {
		return fmt.Errorf("invalid `CategoryID`: %d: %w", r.CategoryID, core_errors.ErrInvalidArgument)
	}

	if r.CityID <= 0 {
		return fmt.Errorf("invalid `CityID`: %d: %w", r.CityID, core_errors.ErrInvalidArgument)
	}

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

	if r.ExpireAt.IsZero() {
		return fmt.Errorf("invalid `ExpireAt`: cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

type UpdateOfferRequest struct {
	Name        *string    `json:"name" validate:"omitempty,min=1,max=100"`
	Description *string    `json:"description" validate:"omitempty,min=1,max=1000"`
	ExpireAt    *time.Time `json:"expire_at" validate:"omitempty"`
}

func (r *UpdateOfferRequest) Validate() error {
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

	if r.ExpireAt != nil && r.ExpireAt.IsZero() {
		return fmt.Errorf("invalid `ExpireAt`: cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

type OfferResponse struct {
	ID          int       `json:"id"`
	PartnerID   int       `json:"partner_id"`
	CategoryID  int       `json:"category_id"`
	CityID      int       `json:"city_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ExpireAt    time.Time `json:"expire_at"`
}