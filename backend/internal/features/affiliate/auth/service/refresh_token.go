package auth_service

import (
	"context"
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_security "github.com/0ScPro0/affiliate-system/internal/core/security"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

// RefreshToken verifies the refresh token, generates new access and refresh tokens,
// updates the refresh token in the database, and returns the new tokens.
func (s *AuthService) RefreshToken(
	ctx context.Context,
	req core_transport_dto.RefreshTokenRequest,
) (core_transport_dto.RefreshTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", err)
	}

	// Verify the refresh token JWT
	claims, err := core_security.VerifyToken(s.cfg, req.RefreshToken)
	if err != nil {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", core_errors.ErrUnauthorized)
	}

	// Extract user ID from claims
	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: invalid token claims: %w", core_errors.ErrUnauthorized)
	}
	userID := int(userIDFloat)

	// Get user from database to verify the stored refresh token
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", core_errors.ErrUnauthorized)
	}

	// Verify that the stored refresh token matches the provided one
	if user.RefreshToken == nil || *user.RefreshToken != req.RefreshToken {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", core_errors.ErrUnauthorized)
	}

	// Verify that the refresh token has not expired
	if user.RefreshTokenExpiresAt != nil && time.Now().UTC().After(*user.RefreshTokenExpiresAt) {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: token expired: %w", core_errors.ErrUnauthorized)
	}

	// Generate new tokens
	sub := map[string]any{
		"id":    user.ID,
		"email": user.Email,
		"role":  "user",
	}

	accessToken, err := core_security.CreateAccessToken(s.cfg, sub)
	if err != nil {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", err)
	}

	newRefreshToken, err := core_security.CreateRefreshToken(s.cfg, sub)
	if err != nil {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", err)
	}

	now := time.Now().UTC()
	accessExpiresAt := now.Add(time.Duration(s.cfg.Security.AccessTokenExpireMinutes) * time.Minute)
	refreshExpiresAt := now.Add(time.Duration(s.cfg.Security.RefreshTokenExpireDays) * time.Hour * 24)

	// Update refresh token in the database
	err = s.userRepository.UpdateUserRefreshToken(ctx, user.ID, &newRefreshToken, &refreshExpiresAt)
	if err != nil {
		return core_transport_dto.RefreshTokenResponse{}, fmt.Errorf("refresh_token: %w", err)
	}

	return core_transport_dto.RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessExpiresAt,
	}, nil
}