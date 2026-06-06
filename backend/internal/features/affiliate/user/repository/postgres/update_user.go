package user_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int,
	user core_transport_dto.UpdateUserRequest,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		UPDATE affiliate_system.users
		SET username = COALESCE(NULLIF($2, ''), username),
			email = COALESCE(NULLIF($3, ''), email),
			password_hash = COALESCE(NULLIF($4, ''), password_hash),
			is_admin = COALESCE(NULLIF($5, ''), is_admin),
		WHERE id = $1
		RETURNING id, username, email, password_hash, is_admin, created_at,
		          refresh_token, refresh_token_expires_at
	`

	row := r.pool.QueryRow(ctx, query, id, user.UserName, user.Email, user.PasswordHash, user.IsAdmin)

	var userModel core_database_models.UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.UserName,
		&userModel.Email,
		&userModel.PasswordHash,
		&userModel.IsAdmin,
		&userModel.CreatedAt,
		&userModel.RefreshToken,
		&userModel.RefreshTokenExpiresAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id %d not found", id)
		}
		return domain.User{}, fmt.Errorf("update user error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.UserName,
		userModel.Email,
		userModel.PasswordHash,
		userModel.IsAdmin,
		userModel.CreatedAt,
	)

	if userModel.RefreshToken != nil {
		userDomain.WithTokens("", userModel.RefreshToken, userModel.RefreshTokenExpiresAt)
	}

	return userDomain, nil
}