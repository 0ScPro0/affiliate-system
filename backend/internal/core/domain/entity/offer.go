package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
)

type Offer struct {
	ID          int
	PartnerID   int
	CategoryID  int
	CityID      int
	Name        string
	Description *string
	CreatedAt   time.Time
	ExpireAt 	time.Time
}

func NewOffer(
	id int,
	partnerId int,
	categoryId int,
	cityId int,
	name string,
	description *string,
	createdAt time.Time,
	expireAt time.Time,
) Offer {
	return Offer{
		ID: id,
		PartnerID: partnerId,
		CategoryID: categoryId,
		CityID: cityId,
		Name: name,
		Description: description,
		CreatedAt: createdAt,
		ExpireAt: expireAt,
	}
}

func (o *Offer) Validate() error {

	// ID's validation
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

	// Name validation
	nameLength := len([]rune(o.Name))
	if nameLength < 1 || nameLength > 100 {
		return fmt.Errorf(
			"Invalid `Name` len: %d: %w", 
			nameLength, 
			core_errors.ErrInvalidArgument,
		)
	}

	// Description validation
	if o.Description != nil {
		descriptionLength := len([]rune(*o.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf(
				"Invalid `Description` len: %d: %w", 
				descriptionLength, 
				core_errors.ErrInvalidArgument,
			)
		}
	}

	// CreatedAt validation
    if o.CreatedAt.IsZero() {
        return fmt.Errorf("created_at cannot be zero: %w", core_errors.ErrInvalidArgument)
    }

	// ExpireAt validation
	if o.ExpireAt.IsZero() {
        return fmt.Errorf("expire_at cannot be zero: %w", core_errors.ErrInvalidArgument)
    }

	if o.ExpireAt.Before(o.CreatedAt) || o.ExpireAt.Equal(o.CreatedAt) {
		return fmt.Errorf("`expire_at` must be after created_at")
	}

	if o.ExpireAt.Before(time.Now()) {
        return fmt.Errorf("offer already expired (expire_at: %v): %w", o.ExpireAt, core_errors.ErrInvalidArgument)
    }

	return nil
}

// Utils
func (o *Offer) IsExpired() bool {
    return time.Now().After(o.ExpireAt)
}

func (o *Offer) IsActive() bool {
    return !o.IsExpired()
}

func (o *Offer) DaysUntilExpiration() int {
    if o.IsExpired() {
        return 0
    }
    remaining := time.Until(o.ExpireAt)
    return int(remaining.Hours() / 24)
}