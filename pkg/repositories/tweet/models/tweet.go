package models

import (
	"time"
	"twitter-uala/pkg/domain/tweet/entities"

	"github.com/google/uuid"
)

type Tweet struct {
	TweetID   string    `bson:"tweet_id"`
	Content   string    `bson:"content"`
	UserID    string    `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
}

func FromDomain(tweet entities.Tweet) Tweet {
	id := tweet.TweetID
	if id == "" {
		id = uuid.New().String()
	}
	return Tweet{
		TweetID: id,
		Content: tweet.Content,
		UserID:  tweet.UserID,
	}
}

func (t *Tweet) ToDomain() *entities.Tweet {
	return &entities.Tweet{
		TweetID: t.TweetID,
		Content: t.Content,
		UserID:  t.UserID,
	}
}

func ModelsToDomain(tweetModels []*Tweet) []*entities.Tweet {
	followEntities := make([]*entities.Tweet, 0, len(tweetModels))
	for _, t := range tweetModels {
		followEntities = append(followEntities, t.ToDomain())
	}

	return followEntities
}
