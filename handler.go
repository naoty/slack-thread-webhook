package main

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

func newHandler(client *slack.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		params := req.Context().Value(paramsKey).(map[string]string)
		for key, value := range params {
			fmt.Printf("key: %v, value: %v\n", key, value)
		}

		// TODO: send a message to Slack
	})
}
