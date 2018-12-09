package main

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

type handler struct {
	slack *slack.Client
}

func (handler handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requestParams := req.Context().Value(paramsKey).(map[string]string)
	for key, value := range requestParams {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	messageParams := slack.NewPostMessageParameters()
	handler.slack.PostMessage("general", "Hello", messageParams)
}
