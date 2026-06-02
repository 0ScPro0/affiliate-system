package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
)

type Category struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}

func NewCategory(
	id int,
	name string,
	description *string,
	createdAt time.Time,
) Category {
	return Category{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func (c *Category) Validate() error {
	// ID validation
	if c.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", c.ID, core_errors.ErrInvalidArgument)
	}

	// Name validation
	nameLength := len([]rune(c.Name))
	if nameLength < 1 || nameLength > 100 {
		return fmt.Errorf(
			"invalid `Name` len: %d: %w",
			nameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	// Description validation
	if c.Description != nil {
		descriptionLength := len([]rune(*c.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf(
				"invalid `Description` len: %d: %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	// CreatedAt validation
	if c.CreatedAt.IsZero() {
		return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}