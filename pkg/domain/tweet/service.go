package tweet

import (
	"context"
	"twitter-uala/pkg/configs"
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/domain/tweet/entities"
)

type Service interface {
	CreateTweet(ctx context.Context, content, userID string) (*entities.Tweet, error)
	ValidateContentLenght(content string) error
	GetTweetsInBatch(ctx context.Context, tweetIDs []string) ([]*entities.Tweet, error)
}

type service struct {
	configs         configs.Config
	tweetRepository repositories.Tweet
}

func NewService(configs configs.Config, tweetRepo repositories.Tweet) Service {
	return &service{
		configs:         configs,
		tweetRepository: tweetRepo,
	}
}
