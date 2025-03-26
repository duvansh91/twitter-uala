package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"twitter-uala/cmd/rest"
	"twitter-uala/pkg/configs"
	"twitter-uala/pkg/dependencies"
	"twitter-uala/pkg/services/rest/handlers/reader"
	"twitter-uala/pkg/services/rest/handlers/writer"
)

func main() {
	// Get configs from yml configuration file based on the environment.
	env := os.Getenv("ENVIRONMENT")
	conf := configs.GetConfigsFromYml(env)

	// Create repositories and external dependencies.
	repos := dependencies.CreateRepositoryDependecies(conf)

	// Create domain services.
	domains := dependencies.CreateDomainDependencies(conf, repos)

	// Create handlers for reading and writing based in HTTP methods.
	writer := writer.NewHandler(domains.UserService)
	reader := reader.NewHandler(domains.UserService)

	// Create routes for each use case of the api.
	basePath := "/twitter-uala/user"
	publishTweetRoute := rest.Route{
		Path:    rest.MergePath(basePath, "{user_id}/publish"),
		Method:  http.MethodPost,
		Handler: writer.PublishTweet,
	}
	followRoute := rest.Route{
		Path:    rest.MergePath(basePath, "{user_id}/follow"),
		Method:  http.MethodPost,
		Handler: writer.Follow,
	}
	timelineRoute := rest.Route{
		Path:    rest.MergePath(basePath, "{user_id}/timeline"),
		Method:  http.MethodGet,
		Handler: reader.GetTimeline,
	}
	routes := []rest.Route{
		publishTweetRoute,
		followRoute,
		timelineRoute,
	}

	// Initiate a web server with all the routes.
	server := rest.NewSever(routes)
	err := http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), server.Router())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server running on port %s", conf.ServerPort)
}
