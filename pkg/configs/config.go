package configs

type (
	Config struct {
		MongoDBConfig MongoDBConfig `yaml:"mongodb_config"`
		CacheConfig   CacheConfig   `yaml:"cache_config"`
	}
	MongoDBConfig struct {
		Uri          string `yaml:"uri"`
		DatabaseName string `yaml:"database_name"`
		// Timeout is represented in seconds
		Timeout int `yaml:"timeout_seconds"`
	}
	CacheConfig struct {
		// Duration is represented in seconds
		Duration int `yaml:"duration_seconds"`
	}
	Collections struct {
		Users   string `yaml:"users_collection"`
		Tweets  string `yaml:"tweets_collection"`
		Follows string `yaml:"follows_collection"`
	}
)
