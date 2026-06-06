package user_service

import (
	"context"
	"time"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

type UserService struct {
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

func NewUserService(
	userRepository UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}