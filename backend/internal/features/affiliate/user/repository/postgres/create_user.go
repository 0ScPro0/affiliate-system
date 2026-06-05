package user_postgres_repository

import (
	"context"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user core_transport_dto.CreateUserRequest,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		INSERT INTO affiliate_system.users (username, email, password_hash, is_admin)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, password_hash, is_admin, created_at,
		          refresh_token, refresh_token_expires_at
	`

	row := r.pool.QueryRow(ctx, query, user.UserName, user.Email, user.PasswordHash, user.IsAdmin)

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
		return domain.User{}, fmt.Errorf("create user error: %w", err)
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