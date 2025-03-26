package models

import (
	"twitter-uala/pkg/domain/follow/entities"
)

type Follow struct {
	FollowerID string `bson:"follower_id"`
	FollowedID string `bson:"followed_id"`
}

func (f *Follow) ToDomain() *entities.Follow {
	return &entities.Follow{
		FollowerID: f.FollowerID,
		FollowedID: f.FollowedID,
	}
}

func ModelsToDomain(followModels []*Follow) []*entities.Follow {
	followEntities := make([]*entities.Follow, 0, len(followModels))
	for _, f := range followModels {
		followEntities = append(followEntities, f.ToDomain())
	}

	return followEntities
}

func FromDomain(follow entities.Follow) Follow {
	return Follow{
		FollowerID: follow.FollowerID,
		FollowedID: follow.FollowedID,
	}
}
