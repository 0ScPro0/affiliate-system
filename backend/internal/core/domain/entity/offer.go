package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type Offer struct {
	ID          int
	PartnerID   int
	CategoryID  int
	CityID      int
	Name        string
	Description *string
	CreatedAt   time.Time
	ExpireAt    time.Time
}

func NewOffer(
	id int,
	partnerID int,
	categoryID int,
	cityID int,
	name string,
	description *string,
	createdAt time.Time,
	expireAt time.Time,
) Offer {
	return Offer{
		ID:          id,
		PartnerID:   partnerID,
		CategoryID:  categoryID,
		CityID:      cityID,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
		ExpireAt:    expireAt,
	}
}

func (o *Offer) Validate() error {
	if o.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", o.ID, core_errors.ErrInvalidArgument)
	}

	if o.PartnerID <= 0 {
		return fmt.Errorf("invalid `PartnerID`: %d: %w", o.PartnerID, core_errors.ErrInvalidArgument)
	}

	if o.CategoryID <= 0 {
		return fmt.Errorf("invalid `CategoryID`: %d: %w", o.CategoryID, core_errors.ErrInvalidArgument)
	}

	if o.CityID <= 0 {
		return fmt.Errorf("invalid `CityID`: %d: %w", o.CityID, core_errors.ErrInvalidArgument)
	}

	if !core_utils.ValidateStringLen(o.Name, 1, 100) {
		return fmt.Errorf(
			"invalid `Name` length: %d (must be 1-100): %w",
			core_utils.GetStringLen(o.Name),
			core_errors.ErrInvalidArgument,
		)
	}

	if o.Description != nil {
		if !core_utils.ValidateStringLen(*o.Description, 1, 1000) {
			return fmt.Errorf(
				"invalid `Description` length: %d (must be 1-1000): %w",
				core_utils.GetStringLen(*o.Description),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if o.CreatedAt.IsZero() {
		return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	if o.ExpireAt.IsZero() {
		return fmt.Errorf("expire_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}