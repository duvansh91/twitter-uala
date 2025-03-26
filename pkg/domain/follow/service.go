package follow

import (
	"context"
	"twitter-uala/pkg/domain/follow/entities"
	"twitter-uala/pkg/domain/repositories"
)

type Service interface {
	CreateFollow(ctx context.Context, follower, followed string) error
	GetFollowers(ctx context.Context, userID string) ([]*entities.Follow, error)
}

type service struct {
	followRepository repositories.Follow
}

func NewService(followRepo repositories.Follow) Service {
	return &service{
		followRepository: followRepo,
	}
}
