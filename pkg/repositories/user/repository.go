package user

import (
	"twitter-uala/pkg/domain/repositories"
	"twitter-uala/pkg/repositories/external/cache"
	"twitter-uala/pkg/repositories/external/mongodb"
)

type repository struct {
	database mongodb.Repository
	cache    cache.Repository
}

func NewRepository(database mongodb.Repository, cache cache.Repository) repositories.User {
	return &repository{
		database: database,
		cache:    cache,
	}
}
