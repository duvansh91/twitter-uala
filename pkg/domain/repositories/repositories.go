package repositories

import (
	"context"
	followEntities "twitter-uala/pkg/domain/follow/entities"
	timelineEntities "twitter-uala/pkg/domain/timeline/entities"
	"twitter-uala/pkg/domain/tweet/entities"
	userEntities "twitter-uala/pkg/domain/user/entities"
)

type (
	User interface {
		CreateUser(ctx context.Context, user userEntities.User) (*userEntities.User, error)
		FindUserByID(ctx context.Context, userID string) (*userEntities.User, error)
	}

	Follow interface {
		CreateFollow(ctx context.Context, follow followEntities.Follow) error
		GetFollowers(ctx context.Context, userID string) ([]*followEntities.Follow, error)
	}

	Timeline interface {
		CreateTimeline(ctx context.Context, timeline timelineEntities.Timeline) error
		CreateTimelineInCache(ctx context.Context, timeline timelineEntities.Timeline) error
		GetLatestsTimeline(ctx context.Context, userID string) ([]*timelineEntities.Timeline, error)
		GetLatestsTimelineFromCache(ctx context.Context, userID string) ([]*timelineEntities.Timeline, error)
	}

	Tweet interface {
		CreateTweet(ctx context.Context, tweet entities.Tweet) (*entities.Tweet, error)
		CreateTweetInCache(ctx context.Context, tweet entities.Tweet) error
		GetTweetsInBatch(ctx context.Context, tweetIDs []string) ([]*entities.Tweet, error)
		GetTweetsInBatchFromCache(ctx context.Context, tweetIDs []string) ([]*entities.Tweet, error)
	}
)
