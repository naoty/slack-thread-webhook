package main

import (
	"fmt"
	"os"

	"github.com/naoty/slack-thread-webhook/datastore"

	"github.com/nlopes/slack"
)

func main() {
	redis := datastore.NewRedis(os.Getenv("REDIS_HOST"))
	if err := redis.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to redis: %v\n", err)
		os.Exit(1)
	}

	client := slack.New(os.Getenv("SLACK_TOKEN"))
	handler := handler{datastore: redis, slack: client}

	router := &router{}
	router.post("/hooks/(?P<id>\\d+)", handler)

	cli := cli{outStream: os.Stdout, errStream: os.Stderr, router: router}
	status := cli.run(os.Args)
	os.Exit(status)
}
