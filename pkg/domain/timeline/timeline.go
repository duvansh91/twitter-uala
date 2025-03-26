package timeline

import (
	"context"
	"twitter-uala/pkg/domain/timeline/entities"
)

func (s *service) CreateTimeline(ctx context.Context, tweetID, userID string) error {
	followers, err := s.followService.GetFollowers(ctx, userID)
	if err != nil {
		return nil
	}

	if followers == nil || len(followers) == 0 {
		return nil
	}

	for _, f := range followers {
		timeline := entities.Timeline{
			UserID:  f.FollowerID,
			TweetID: tweetID,
		}
		err := s.timelineRepository.CreateTimeline(ctx, timeline)
		if err != nil {
			return err
		}

		err = s.timelineRepository.CreateTimelineInCache(ctx, timeline)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetLatestsTimeline(ctx context.Context, userID string) ([]*entities.Timeline, error) {
	timelines, err := s.timelineRepository.GetLatestsTimelineFromCache(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(timelines) == 0 {
		return s.timelineRepository.GetLatestsTimeline(ctx, userID)
	}
	return timelines, nil
}

func (s *service) GetTweetIDsFromTimelines(timelines []*entities.Timeline) []string {
	var tweetIDs []string
	for _, t := range timelines {
		tweetIDs = append(tweetIDs, t.TweetID)
	}

	return tweetIDs
}
