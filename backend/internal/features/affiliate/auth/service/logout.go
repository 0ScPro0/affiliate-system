package auth_service

import (
	"context"
	"fmt"
)

// Logout clears the refresh token for the authenticated user.
func (s *AuthService) Logout(
	ctx context.Context,
	userID int,
) error {
	err := s.userRepository.UpdateUserRefreshToken(ctx, userID, nil, nil)
	if err != nil {
		return fmt.Errorf("logout: %w", err)
	}

	return nil
}