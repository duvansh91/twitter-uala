package dependencies

import (
	"time"
	"twitter-uala/pkg/configs"
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/repositories/external/cache"
	"twitter-uala/pkg/repositories/external/mongodb"
	"twitter-uala/pkg/repositories/follow"
	"twitter-uala/pkg/repositories/timeline"
	"twitter-uala/pkg/repositories/tweet"
	"twitter-uala/pkg/repositories/user"
)

type Repositories struct {
	UserRepo     repositories.User
	FollowRepo   repositories.Follow
	TweetRepo    repositories.Tweet
	TimelineRepo repositories.Timeline
}

func CreateRepositoryDependecies(conf configs.Config) Repositories {
	userDBRepo := mongodb.NewRepository(&conf.MongoDBConfig, conf.Collections.Users)
	followDBRepo := mongodb.NewRepository(&conf.MongoDBConfig, conf.Collections.Follows)
	tweetDBRepo := mongodb.NewRepository(&conf.MongoDBConfig, conf.Collections.Tweets)
	timelineDBRepo := mongodb.NewRepository(&conf.MongoDBConfig, conf.Collections.Timelines)

	cacheRepo := cache.NewRepository(time.Duration(conf.CacheConfig.Duration) * time.Hour)

	return Repositories{
		UserRepo:     user.NewRepository(userDBRepo, cacheRepo),
		FollowRepo:   follow.NewRepository(followDBRepo, cacheRepo),
		TweetRepo:    tweet.NewRepository(tweetDBRepo, cacheRepo),
		TimelineRepo: timeline.NewRepository(timelineDBRepo, cacheRepo, conf),
	}
}
