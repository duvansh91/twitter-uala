package configs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		MongoDBConfig      MongoDBConfig `yaml:"mongodb_config"`
		CacheConfig        CacheConfig   `yaml:"cache_config"`
		MaxTweetLength     int           `yaml:"max_tweet_length"`
		LatestsTweetsLimit int64         `yaml:"latests_tweets_limit"`
		Collections        Collections   `yaml:"collections"`
		ServerPort         string        `yaml:"server_port"`
	}
	MongoDBConfig struct {
		Uri          string `yaml:"uri"`
		DatabaseName string `yaml:"database_name"`
		// Timeout is represented in seconds
		Timeout int `yaml:"timeout_seconds"`
	}
	CacheConfig struct {
		// Duration is represented in hours
		Duration int `yaml:"duration_hours"`
	}
	Collections struct {
		Users     string `yaml:"users_collection"`
		Tweets    string `yaml:"tweets_collection"`
		Follows   string `yaml:"follows_collection"`
		Timelines string `yaml:"timelines_collection"`
	}
)

func GetConfigsFromYml(env string) Config {
	if env == "" {
		env = "local"
	}
	basePath, _ := os.Getwd()
	path := fmt.Sprintf("%s/config/%s.yml", basePath, strings.ToLower(env))

	file, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprint("error reading yml file ", err.Error()))
	}
	var conf Config
	ParseConfigsFromYml(file, &conf)

	return conf
}

func ParseConfigsFromYml(file []byte, conf *Config) {
	err := yaml.Unmarshal(file, conf)
	if err != nil {
		panic("invalid yml schema")
	}
}

// /twitter-uala/cmd/config/local.yml
