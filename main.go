package main

import (
	"fmt"
	"os"

	"github.com/naoty/slack-thread-webhook/datastore"
	"github.com/naoty/slack-thread-webhook/handler"

	"github.com/nlopes/slack"
)

func main() {
	redis := &datastore.Redis{URL: os.Getenv("REDIS_URL")}
	if err := redis.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to redis: %v\n", err)
		os.Exit(1)
	}

	client := slack.New(os.Getenv("SLACK_TOKEN"))
	channels, err := client.GetChannels(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get channels: %v\n", err)
		os.Exit(1)
	}

	channelID := ""
	for _, channel := range channels {
		if channel.Name == os.Getenv("SLACK_CHANNEL") {
			channelID = channel.ID
			break
		}
	}

	if channelID == "" {
		fmt.Fprintf(os.Stderr, "channel not found: %v\n", os.Getenv("SLACK_CHANNEL"))
		os.Exit(1)
	}

	post := handler.Post{Channel: channelID, Datastore: redis, Slack: client}
	put := handler.Put{Channel: channelID, Datastore: redis, Slack: client}

	router := &handler.Router{}
	router.POST("/hooks/(?P<id>\\S+)", post)
	router.PUT("/hooks/(?P<id>\\S+)", put)

	cli := cli{outStream: os.Stdout, errStream: os.Stderr, router: router}
	status := cli.run(os.Args)
	os.Exit(status)
}
