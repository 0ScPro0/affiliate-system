package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type User struct {
	ID           int
	UserName     *string
	Email        string
	PasswordHash string
	IsAdmin      bool
	CreatedAt    time.Time
}

func NewUser(
	id int,
	username *string,
	email string,
	passwordHash string,
	isAdmin bool,
	createdAt time.Time,
) User {
	return User{
		ID:           id,
		UserName:     username,
		Email:        email,
		PasswordHash: passwordHash,
		IsAdmin:      isAdmin,
		CreatedAt:    createdAt,
	}
}

func (u *User) Validate() error {
	if u.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", u.ID, core_errors.ErrInvalidArgument)
	}

	if u.UserName != nil {
		if !core_utils.ValidateStringLen(*u.UserName, 1, 50) {
			return fmt.Errorf(
				"invalid `UserName` length: %d (must be 1-50): %w",
				core_utils.GetStringLen(*u.UserName),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if err := u.validateEmail(); err != nil {
		return err
	}

	if !core_utils.ValidateStringLen(u.PasswordHash, 1, 255) {
		return fmt.Errorf(
			"invalid `PasswordHash` length: %d (must be 1-255): %w",
			core_utils.GetStringLen(u.PasswordHash),
			core_errors.ErrInvalidArgument,
		)
	}

	if u.CreatedAt.IsZero() {
		return fmt.Errorf("invalid `CreatedAt`: cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (u *User) validateEmail() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !core_utils.ValidateStringLen(u.Email, 1, 100) {
		return fmt.Errorf(
			"invalid `Email` length: %d (must be 1-100): %w",
			core_utils.GetStringLen(u.Email),
			core_errors.ErrInvalidArgument,
		)
	}

	if !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("invalid `Email` format: %s: %w", u.Email, core_errors.ErrInvalidArgument)
	}

	return nil
}
