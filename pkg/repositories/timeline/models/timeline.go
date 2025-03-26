package models

import (
	"time"
	"twitter-uala/pkg/domain/timeline/entities"
)

type TimelineCache struct {
	TweetIDs []string `json:"tweet_ids"`
}

type Timeline struct {
	TimelineID string    `bson:"_id"`
	UserID     string    `bson:"user_id"`
	TweetID    string    `bson:"tweet_id"`
	CreatedAt  time.Time `bson:"created_at"`
}

type TimelineToInsert struct {
	UserID    string    `bson:"user_id"`
	TweetID   string    `bson:"tweet_id"`
	CreatedAt time.Time `bson:"created_at"`
}

func FromDomain(timline entities.Timeline) TimelineToInsert {
	return TimelineToInsert{
		UserID:  timline.UserID,
		TweetID: timline.TweetID,
	}
}

func (t *Timeline) ToDomain() *entities.Timeline {
	return &entities.Timeline{
		TimelineID: t.TimelineID,
		UserID:     t.UserID,
		TweetID:    t.TweetID,
	}
}

func ModelsToDomain(timelineModels []*Timeline) []*entities.Timeline {
	timelineEntities := make([]*entities.Timeline, 0, len(timelineModels))
	for _, t := range timelineModels {
		timelineEntities = append(timelineEntities, t.ToDomain())
	}

	return timelineEntities
}

func (tc *TimelineCache) ToDomain() []*entities.Timeline {
	timelineEntities := make([]*entities.Timeline, 0, len(tc.TweetIDs))
	for _, id := range tc.TweetIDs {
		timelineEntities = append(timelineEntities, &entities.Timeline{TweetID: id})
	}

	return timelineEntities
}
