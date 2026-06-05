package user_postgres_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *UserRepository) UpdateUserRefreshToken(
	ctx context.Context,
	id int,
	refreshToken *string,
	refreshTokenExpiresAt *time.Time,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		UPDATE affiliate_system.users
		SET refresh_token = $2,
		    refresh_token_expires_at = $3
		WHERE id = $1
	`

	_, err := r.pool.Exec(ctx, query, id, refreshToken, refreshTokenExpiresAt)
	if err != nil {
		return fmt.Errorf("update user refresh token error: %w", err)
	}

	return nil
}