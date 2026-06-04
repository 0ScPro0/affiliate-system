package user_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_database_models "github.com/0ScPro0/affiliate-system/internal/core/database/models"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		SELECT id, username, email, password_hash, is_admin, created_at
		FROM affiliate_system.users
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel core_database_models.UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.UserName,
		&userModel.Email,
		&userModel.PasswordHash,
		&userModel.IsAdmin,
		&userModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id %d not found", id)
		}
		return domain.User{}, fmt.Errorf("get user by id error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.UserName,
		userModel.Email,
		userModel.PasswordHash,
		userModel.IsAdmin,
		userModel.CreatedAt,
	)

	return userDomain, nil
}

func (r *UserRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeout())
	defer cancel()

	query := `
		SELECT id, username, email, password_hash, is_admin, created_at
		FROM affiliate_system.users
		WHERE email = $1
	`

	row := r.pool.QueryRow(ctx, query, email)

	var userModel core_database_models.UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.UserName,
		&userModel.Email,
		&userModel.PasswordHash,
		&userModel.IsAdmin,
		&userModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with email %s not found", email)
		}
		return domain.User{}, fmt.Errorf("get user by email error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.UserName,
		userModel.Email,
		userModel.PasswordHash,
		userModel.IsAdmin,
		userModel.CreatedAt,
	)

	return userDomain, nil
}