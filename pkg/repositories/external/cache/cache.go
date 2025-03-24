package cache

import "github.com/karlseguin/ccache/v2"

func (r *repository) Get(key string) *ccache.Item {
	return r.cache.Get(key)
}

func (r *repository) Set(key string, item []byte) {
	r.cache.Set(key, item, r.durationTime)
}
