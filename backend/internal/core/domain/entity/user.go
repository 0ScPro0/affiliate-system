package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
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
	// ID validation
	if u.ID < 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", u.ID, core_errors.ErrInvalidArgument)
	}

	// UserName validation (optional)
	if u.UserName != nil {
		usernameLength := len([]rune(*u.UserName))
		if usernameLength < 1 || usernameLength > 50 {
			return fmt.Errorf(
				"invalid `UserName` length: %d (must be 1-50): %w",
				usernameLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	// Email validation
	if err := u.validateEmail(); err != nil {
		return err
	}

	// PasswordHash validation
	if len(u.PasswordHash) == 0 {
		return fmt.Errorf("invalid `PasswordHash`: cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if len(u.PasswordHash) > 255 {
		return fmt.Errorf("invalid `PasswordHash` length: %d (max 255): %w",
			len(u.PasswordHash),
			core_errors.ErrInvalidArgument,
		)
	}

	// CreatedAt validation
	if u.CreatedAt.IsZero() {
		return fmt.Errorf("invalid `CreatedAt`: cannot be zero: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (u *User) validateEmail() error {
	// Basic email regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if len(u.Email) < 1 || len(u.Email) > 100 {
		return fmt.Errorf(
			"invalid `Email` length: %d (must be 1-100): %w",
			len(u.Email),
			core_errors.ErrInvalidArgument,
		)
	}

	if !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("invalid `Email` format: %s: %w", u.Email, core_errors.ErrInvalidArgument)
	}

	return nil
}
