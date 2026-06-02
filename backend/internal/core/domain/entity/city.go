package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
)

type City struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func NewCity(id int, name string, createdAt time.Time) City {
	return City{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}

func (c *City) Validate() error {
	if c.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", c.ID, core_errors.ErrInvalidArgument)
	}

	nameLength := len([]rune(c.Name))
	if nameLength < 1 || nameLength > 50 {
		return fmt.Errorf(
			"invalid `Name` len: %d: %w",
			nameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if c.CreatedAt.IsZero() {
		return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}