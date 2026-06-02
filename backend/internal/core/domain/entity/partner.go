package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
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
	// ID validation
	if p.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", p.ID, core_errors.ErrInvalidArgument)
	}

	// Name validation
	nameLength := len([]rune(p.Name))
	if nameLength < 1 || nameLength > 100 {
		return fmt.Errorf(
			"invalid `Name` len: %d: %w",
			nameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	// Description validation
	if p.Description != nil {
		descriptionLength := len([]rune(*p.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf(
				"invalid `Description` len: %d: %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	// CreatedAt validation
	if p.CreatedAt.IsZero() {
		return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}