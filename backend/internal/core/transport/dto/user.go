package core_transport_dto

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

type CreateUserRequest struct {
	UserName     *string `json:"username" validate:"omitempty,min=1,max=50"`
	Email        string  `json:"email" validate:"required,email,max=100"`
	PasswordHash string  `json:"password_hash" validate:"required,min=1,max=255"`
	IsAdmin      bool    `json:"is_admin"`
}

func (r *CreateUserRequest) Validate() error {
	if r.UserName != nil {
		if !core_utils.ValidateStringLen(*r.UserName, 1, 50) {
			return fmt.Errorf(
				"invalid `UserName` length: %d (must be 1-50): %w",
				core_utils.GetStringLen(*r.UserName),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if err := validateEmail(r.Email); err != nil {
		return err
	}

	if !core_utils.ValidateStringLen(r.PasswordHash, 1, 255) {
		return fmt.Errorf(
			"invalid `PasswordHash` length: %d (must be 1-255): %w",
			core_utils.GetStringLen(r.PasswordHash),
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type UpdateUserRequest struct {
	ID           int     `json:"id" validate:"required"`
	UserName     *string `json:"username" validate:"omitempty,min=1,max=50"`
	Email        *string `json:"email" validate:"omitempty,email,max=100"`
	PasswordHash *string `json:"password_hash" validate:"omitempty,min=1,max=255"`
	IsAdmin      *bool   `json:"is_admin"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.ID <= 0 {
		return fmt.Errorf("invalid `ID`: %d: %w", r.ID, core_errors.ErrInvalidArgument)
	}

	if r.UserName != nil {
		if !core_utils.ValidateStringLen(*r.UserName, 1, 50) {
			return fmt.Errorf(
				"invalid `UserName` length: %d (must be 1-50): %w",
				core_utils.GetStringLen(*r.UserName),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Email != nil {
		if err := validateEmail(*r.Email); err != nil {
			return err
		}
	}

	if r.PasswordHash != nil {
		if !core_utils.ValidateStringLen(*r.PasswordHash, 1, 255) {
			return fmt.Errorf(
				"invalid `PasswordHash` length: %d (must be 1-255): %w",
				core_utils.GetStringLen(*r.PasswordHash),
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type UserResponse struct {
	ID        int       `json:"id"`
	UserName  *string   `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

func validateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !core_utils.ValidateStringLen(email, 1, 100) {
		return fmt.Errorf(
			"invalid `Email` length: %d (must be 1-100): %w",
			core_utils.GetStringLen(email),
			core_errors.ErrInvalidArgument,
		)
	}

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid `Email` format: %s: %w", email, core_errors.ErrInvalidArgument)
	}

	return nil
}