package core_transport_dto

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_utils "github.com/0ScPro0/affiliate-system/internal/core/utils"
)

// ===================== REGISTER =====================
type RegisterRequest struct {
	UserName *string `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func (r *RegisterRequest) Validate() error {
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

	if !core_utils.ValidateStringLen(r.Password, 8, 255) {
		return fmt.Errorf(
			"invalid `Password` length: %d (must be 8-255): %w",
			core_utils.GetStringLen(r.Password),
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type RegisterResponse struct {
	UserResponse
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

// ===================== LOGIN =====================
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if err := validateEmail(r.Email); err != nil {
		return err
	}

	if !core_utils.ValidateStringLen(r.Password, 1, 255) {
		return fmt.Errorf(
			"invalid `Password` length: %d (must be 1-255): %w",
			core_utils.GetStringLen(r.Password),
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type LoginResponse struct {
	UserResponse
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	RefreshToken         string    `json:"refresh_token"`
}

// ===================== REFRESH TOKEN =====================
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r *RefreshTokenRequest) Validate() error {
	if !core_utils.ValidateStringLen(r.RefreshToken, 1, 512) {
		return fmt.Errorf(
			"invalid `RefreshToken` length: %d (must be 1-512): %w",
			core_utils.GetStringLen(r.RefreshToken),
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type RefreshTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}