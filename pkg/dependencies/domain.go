package dependencies

import (
	"twitter-uala/pkg/configs"
	"twitter-uala/pkg/domain/follow"
	"twitter-uala/pkg/domain/timeline"
	"twitter-uala/pkg/domain/tweet"
	"twitter-uala/pkg/domain/user"
)

type Domains struct {
	UserService     user.Service
	FollowService   follow.Service
	TweetService    tweet.Service
	TimelineService timeline.Service
}

func CreateDomainDependencies(conf configs.Config, repositories Repositories) Domains {
	followServ := follow.NewService(repositories.FollowRepo)
	timelineServ := timeline.NewService(repositories.TimelineRepo, followServ)
	tweetServ := tweet.NewService(conf, repositories.TweetRepo)
	userServ := user.NewService(
		repositories.UserRepo,
		followServ,
		timelineServ,
		tweetServ,
	)

	return Domains{
		UserService:     userServ,
		FollowService:   followServ,
		TweetService:    tweetServ,
		TimelineService: timelineServ,
	}
}
