package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
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
	if c.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", c.ID, core_errors.ErrInvalidArgument)
	}

	if !core_utils.ValidateStringLen(c.Name, 1, 100) {
		return fmt.Errorf(
			"invalid `Name` length: %d (must be 1-100): %w",
			core_utils.GetStringLen(c.Name),
			core_errors.ErrInvalidArgument,
		)
	}

	if c.Description != nil {
		if !core_utils.ValidateStringLen(*c.Description, 1, 1000) {
			return fmt.Errorf(
				"invalid `Description` length: %d (must be 1-1000): %w",
				core_utils.GetStringLen(*c.Description),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if c.CreatedAt.IsZero() {
		return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}