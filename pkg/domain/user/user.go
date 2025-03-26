package user

import (
	"context"
	"strings"
	"twitter-uala/pkg/constants"
	tweetEntities "twitter-uala/pkg/domain/tweet/entities"
	"twitter-uala/pkg/domain/user/entities"
)

func (s *service) FindOrCreateUser(ctx context.Context, userID string) (*entities.User, error) {
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		if strings.Contains(err.Error(), constants.UserNotFoundError) {
			return s.userRepository.CreateUser(ctx, entities.User{UserID: userID})
		}
		return nil, err
	}

	return user, nil
}

func (s *service) PublishTweet(ctx context.Context, content, userID string) error {
	err := s.tweetService.ValidateContentLenght(content)
	if err != nil {
		return err
	}

	user, err := s.FindOrCreateUser(ctx, userID)
	if err != nil {
		return err
	}

	tweet, err := s.tweetService.CreateTweet(ctx, content, user.UserID)
	if err != nil {
		return err
	}

	return s.timelineService.CreateTimeline(ctx, tweet.TweetID, user.UserID)
}

func (s *service) Follow(ctx context.Context, followerID, followedID string) error {
	follower, err := s.FindOrCreateUser(ctx, followerID)
	if err != nil {
		return err
	}
	followed, err := s.FindOrCreateUser(ctx, followedID)
	if err != nil {
		return err
	}

	err = s.followService.CreateFollow(ctx, follower.UserID, followed.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetTweetsTimeline(ctx context.Context, userID string) ([]*tweetEntities.Tweet, error) {
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		if strings.Contains(err.Error(), constants.UserNotFoundError) {
			return []*tweetEntities.Tweet{}, nil
		}
		return nil, err
	}

	timelines, err := s.timelineService.GetLatestsTimeline(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	if timelines == nil || len(timelines) == 0 {
		return []*tweetEntities.Tweet{}, nil
	}

	tweets, err := s.tweetService.GetTweetsInBatch(ctx, s.timelineService.GetTweetIDsFromTimelines(timelines))
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
