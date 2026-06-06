package user_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
)

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int,
	user core_transport_dto.UpdateUserRequest,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate update user request: %w", err)
	}

	domainUser, err := s.userRepository.UpdateUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("update user: %w", err)
	}

	return domainUser, nil
}