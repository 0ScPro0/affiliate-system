package auth_service

import (
	"context"
	"time"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

type AuthService struct {
	cfg            *config.Config
	userRepository UserRepository
}

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user core_transport_dto.CreateUserRequest,
	) (domain.User, error)

	GetUserByID(
		ctx context.Context,
		id int,
	) (domain.User, error)

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (domain.User, error)

	UpdateUser(
		ctx context.Context,
		id int,
		user core_transport_dto.UpdateUserRequest,
	) (domain.User, error)

	UpdateUserRefreshToken(
		ctx context.Context,
		id int,
		refreshToken *string,
		refreshTokenExpiresAt *time.Time,
	) error

	DeleteUser(
		ctx context.Context,
		id int,
	) error
}

func NewAuthService(
	cfg *config.Config,
	userRepository UserRepository,
) *AuthService {
	return &AuthService{
		cfg:            cfg,
		userRepository: userRepository,
	}
}