package follow

import (
	"context"
	"fmt"
	"twitter-uala/pkg/domain/follow/entities"
)

func (s *service) CreateFollow(ctx context.Context, follower, followed string) error {
	follow := entities.Follow{
		FollowerID: follower,
		FollowedID: followed,
	}
	err := s.followRepository.CreateFollow(ctx, follow)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetFollowers(ctx context.Context, userID string) ([]*entities.Follow, error) {
	followers, err := s.followRepository.GetFollowers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting followers of user %s: %s", userID, err.Error())
	}

	return followers, nil
}
