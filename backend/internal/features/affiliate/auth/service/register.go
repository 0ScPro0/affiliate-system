package auth_service

import (
	"context"
	"fmt"
	"time"

	core_security "github.com/0ScPro0/affiliate-system/internal/core/security"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

// Register creates a new user, generates access and refresh tokens, and returns them.
func (s *AuthService) Register(
	ctx context.Context,
	req core_transport_dto.RegisterRequest,
) (core_transport_dto.RegisterResponse, error) {
	if err := req.Validate(); err != nil {
		return core_transport_dto.RegisterResponse{}, fmt.Errorf("register: %w", err)
	}

	passwordHash, err := core_security.HashPassword(req.Password)
	if err != nil {
		return core_transport_dto.RegisterResponse{}, fmt.Errorf("register: hash password: %w", err)
	}

	createReq := core_transport_dto.CreateUserRequest{
		UserName:     req.UserName,
		Email:        req.Email,
		PasswordHash: passwordHash,
		IsAdmin:      false,
	}

	user, err := s.userRepository.CreateUser(ctx, createReq)
	if err != nil {
		return core_transport_dto.RegisterResponse{}, fmt.Errorf("register: %w", err)
	}

	sub := map[string]any{
		"sub":      user.ID,
		"is_admin": user.IsAdmin,
	}

	accessToken, err := core_security.CreateAccessToken(s.cfg, sub)
	if err != nil {
		return core_transport_dto.RegisterResponse{}, fmt.Errorf("register: %w", err)
	}

	refreshToken, err := core_security.CreateRefreshToken(s.cfg, sub)
	if err != nil {
		return core_transport_dto.RegisterResponse{}, fmt.Errorf("register: %w", err)
	}

	now := time.Now().UTC()
	accessExpiresAt := now.Add(time.Duration(s.cfg.Security.AccessTokenExpireMinutes) * time.Minute)
	refreshExpiresAt := now.Add(time.Duration(s.cfg.Security.RefreshTokenExpireDays) * time.Hour * 24)

	user.WithTokens(accessToken, &refreshToken, &refreshExpiresAt)

	return core_transport_dto.RegisterResponse{
		UserResponse: core_transport_dto.UserResponse{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			IsAdmin:   user.IsAdmin,
			CreatedAt: user.CreatedAt,
		},
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshExpiresAt,
	}, nil
}