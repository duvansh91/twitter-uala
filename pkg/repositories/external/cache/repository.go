package cache

import (
	"sync"
	"time"

	"github.com/karlseguin/ccache/v2"
)

var (
	once      sync.Once
	initiated *repository
)

type repository struct {
	cache        *ccache.Cache
	durationTime time.Duration
}

type Repository interface {
	Get(key string) *ccache.Item
	Set(key string, item []byte)
}

func NewRepository(duration time.Duration) Repository {
	config := ccache.Configure()
	cache := ccache.New(config)

	// Cache must be initiated only one time.
	once.Do(
		func() {
			initiated = &repository{
				cache:        cache,
				durationTime: duration,
			}
		},
	)

	return initiated
}
