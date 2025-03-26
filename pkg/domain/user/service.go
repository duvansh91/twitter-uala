package user

import (
	"context"
	"twitter-uala/pkg/domain/follow"
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/domain/timeline"
	"twitter-uala/pkg/domain/tweet"
	tweetEntities "twitter-uala/pkg/domain/tweet/entities"
	"twitter-uala/pkg/domain/user/entities"
)

type Service interface {
	FindOrCreateUser(ctx context.Context, userID string) (*entities.User, error)
	PublishTweet(ctx context.Context, content, userID string) error
	Follow(ctx context.Context, followerID, followedID string) error
	GetTweetsTimeline(ctx context.Context, userID string) ([]*tweetEntities.Tweet, error)
}

type service struct {
	userRepository  repositories.User
	followService   follow.Service
	timelineService timeline.Service
	tweetService    tweet.Service
}

func NewService(
	userRepo repositories.User,
	followServ follow.Service,
	timelineServ timeline.Service,
	tweetServ tweet.Service,
) Service {
	return &service{
		userRepository:  userRepo,
		followService:   followServ,
		timelineService: timelineServ,
		tweetService:    tweetServ,
	}
}
