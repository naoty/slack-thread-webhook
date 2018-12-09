package main

import (
	"github.com/naoty/slack-thread-webhook/datastore"
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

type handler struct {
	datastore datastore.Client
	slack *slack.Client
}

func (handler handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	messageParams := slack.NewPostMessageParameters()
	_, ts, err := handler.slack.PostMessage("general", "Hello", messageParams)
	if err != nil {
		message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
		http.Error(w, message, http.StatusInternalServerError)
	}

	requestParams := req.Context().Value(paramsKey).(map[string]string)
	id := requestParams["id"]
	err = handler.datastore.Set(id, ts)
	if err != nil {
		message := fmt.Sprintf("failed to set id: %v\n", err)
		http.Error(w, message, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
