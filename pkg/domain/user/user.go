package user

import (
	"context"
	"twitter-uala/pkg/domain/user/entities"
)

func (s *service) CreateUser(ctx context.Context, user entities.User) error {
	return s.userRepository.CreateUser(ctx, user)
}
