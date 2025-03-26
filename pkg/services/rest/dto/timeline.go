package dto

import "twitter-uala/pkg/domain/tweet/entities"

type TimelineResponse struct {
	Tweets []TweetResponse `json:"tweets"`
}

func NewTimelineResponse(tweets []*entities.Tweet) TimelineResponse {
	tweetsResponse := make([]TweetResponse, 0, len(tweets))
	for _, tEntity := range tweets {
		tweet := TweetResponse{
			UserID:  tEntity.UserID,
			Content: tEntity.Content,
		}
		tweetsResponse = append(tweetsResponse, tweet)
	}

	return TimelineResponse{
		Tweets: tweetsResponse,
	}
}
