package auth_service

import (
	"context"
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_security "github.com/0ScPro0/affiliate-system/internal/core/security"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

// Login authenticates a user by email and password, returns tokens.
func (s *AuthService) Login(
	ctx context.Context,
	req core_transport_dto.LoginRequest,
) (core_transport_dto.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return core_transport_dto.LoginResponse{}, fmt.Errorf("login: %w", err)
	}

	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return core_transport_dto.LoginResponse{}, fmt.Errorf("login: %w", core_errors.ErrUnauthorized)
	}

	if !core_security.VerifyPassword(req.Password, user.PasswordHash) {
		return core_transport_dto.LoginResponse{}, fmt.Errorf("login: %w", core_errors.ErrUnauthorized)
	}

	sub := map[string]any{
		"sub":      user.ID,
		"is_admin": user.IsAdmin,
	}

	accessToken, err := core_security.CreateAccessToken(s.cfg, sub)
	if err != nil {
		return core_transport_dto.LoginResponse{}, fmt.Errorf("login: %w", err)
	}

	refreshToken, err := core_security.CreateRefreshToken(s.cfg, sub)
	if err != nil {
		return core_transport_dto.LoginResponse{}, fmt.Errorf("login: %w", err)
	}

	now := time.Now().UTC()
	accessExpiresAt := now.Add(time.Duration(s.cfg.Security.AccessTokenExpireMinutes) * time.Minute)
	refreshExpiresAt := now.Add(time.Duration(s.cfg.Security.RefreshTokenExpireDays) * time.Hour * 24)

	// Save refresh token in the database
	err = s.userRepository.UpdateUserRefreshToken(ctx, user.ID, &refreshToken, &refreshExpiresAt)
	if err != nil {
		return core_transport_dto.LoginResponse{}, fmt.Errorf("login: %w", err)
	}

	user.WithTokens(accessToken, &refreshToken, &refreshExpiresAt)

	return core_transport_dto.LoginResponse{
		UserResponse: core_transport_dto.UserResponse{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			IsAdmin:   user.IsAdmin,
			CreatedAt: user.CreatedAt,
		},
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessExpiresAt,
		RefreshToken:         refreshToken,
	}, nil
}