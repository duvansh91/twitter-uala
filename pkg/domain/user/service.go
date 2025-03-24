package user

import (
	"context"
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/domain/user/entities"
)

type Service interface {
	CreateUser(ctx context.Context, user entities.User) error
}

type service struct {
	userRepository repositories.User
}

func NewService(userRepo repositories.User) Service {
	return &service{
		userRepository: userRepo,
	}
}
