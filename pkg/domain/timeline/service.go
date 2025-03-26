package timeline

import (
	"context"
	"twitter-uala/pkg/domain/follow"
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/domain/timeline/entities"
)

type Service interface {
	CreateTimeline(ctx context.Context, tweetID, userID string) error
	GetLatestsTimeline(ctx context.Context, userID string) ([]*entities.Timeline, error)
	GetTweetIDsFromTimelines(timelines []*entities.Timeline) []string
}

type service struct {
	timelineRepository repositories.Timeline
	followService      follow.Service
}

func NewService(
	timelineRepo repositories.Timeline,
	followServ follow.Service,
) Service {
	return &service{
		timelineRepository: timelineRepo,
		followService:      followServ,
	}
}
