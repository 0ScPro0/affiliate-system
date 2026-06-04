package user_service

import (
	"context"
	"fmt"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
)

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int,
) (domain.User, error) {
	domainUser, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user by id: %w", err)
	}

	return domainUser, nil
}

func (s *UserService) GetUserByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {
	domainUser, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user by email: %w", err)
	}

	return domainUser, nil
}