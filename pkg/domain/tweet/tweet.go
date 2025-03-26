package tweet

import (
	"context"
	"fmt"
	"twitter-uala/pkg/domain/tweet/entities"
)

func (s *service) CreateTweet(ctx context.Context, content, userID string) (*entities.Tweet, error) {
	tweet := entities.Tweet{
		Content: content,
		UserID:  userID,
	}

	createdTweet, err := s.tweetRepository.CreateTweet(ctx, tweet)
	if err != nil {
		return nil, err
	}

	err = s.tweetRepository.CreateTweetInCache(ctx, *createdTweet)
	if err != nil {
		return nil, err
	}

	return createdTweet, nil
}

func (s *service) ValidateContentLenght(content string) error {
	if len(content) > s.configs.MaxTweetLength {
		return fmt.Errorf(
			"tweet content is %d, must be less than %d",
			len(content),
			s.configs.MaxTweetLength,
		)
	}

	return nil
}

func (s *service) GetTweetsInBatch(ctx context.Context, tweetIDs []string) ([]*entities.Tweet, error) {
	tweets, err := s.tweetRepository.GetTweetsInBatchFromCache(ctx, tweetIDs)
	if err != nil {
		return nil, err
	}

	if len(tweets) == 0 {
		tweets, err = s.tweetRepository.GetTweetsInBatch(ctx, tweetIDs)
		if err != nil {
			return nil, err
		}
	}

	return tweets, nil
}
