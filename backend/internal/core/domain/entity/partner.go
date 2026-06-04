package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type Partner struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}

func NewPartner(
	id int,
	name string,
	description *string,
	createdAt time.Time,
) Partner {
	return Partner{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func (p *Partner) Validate() error {
	if p.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", p.ID, core_errors.ErrInvalidArgument)
	}

	if !core_utils.ValidateStringLen(p.Name, 1, 100) {
		return fmt.Errorf(
			"invalid `Name` length: %d (must be 1-100): %w",
			core_utils.GetStringLen(p.Name),
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Description != nil {
		if !core_utils.ValidateStringLen(*p.Description, 1, 1000) {
			return fmt.Errorf(
				"invalid `Description` length: %d (must be 1-1000): %w",
				core_utils.GetStringLen(*p.Description),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.CreatedAt.IsZero() {
		return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}