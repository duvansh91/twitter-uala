package timeline

import (
	"twitter-uala/pkg/configs"
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/repositories/external/cache"
	"twitter-uala/pkg/repositories/external/mongodb"
)

type repository struct {
	database mongodb.Repository
	cache    cache.Repository
	configs  configs.Config
}

func NewRepository(
	database mongodb.Repository,
	cache cache.Repository,
	configs configs.Config,
) repositories.Timeline {
	return &repository{
		database: database,
		cache:    cache,
		configs:  configs,
	}
}
