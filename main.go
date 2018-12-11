package main

import (
	"fmt"
	"os"

	"github.com/naoty/slack-thread-webhook/datastore"
	"github.com/naoty/slack-thread-webhook/handler"

	"github.com/nlopes/slack"
)

func main() {
	redis := datastore.NewRedis(os.Getenv("REDIS_HOST"))
	if err := redis.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to redis: %v\n", err)
		os.Exit(1)
	}

	client := slack.New(os.Getenv("SLACK_TOKEN"))
	post := handler.Post{Datastore: redis, Slack: client}

	router := &handler.Router{}
	router.POST("/hooks/(?P<id>\\d+)", post)

	cli := cli{outStream: os.Stdout, errStream: os.Stderr, router: router}
	status := cli.run(os.Args)
	os.Exit(status)
}
