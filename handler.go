package main

import (
	"fmt"
	"net/http"

	"github.com/naoty/slack-thread-webhook/datastore"

	"github.com/nlopes/slack"
)

type handler struct {
	datastore datastore.Client
	slack     *slack.Client
}

func (handler handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	messageParams := slack.NewPostMessageParameters()
	requestParams := req.Context().Value(paramsKey).(map[string]string)
	id := requestParams["id"]
	value, _ := handler.datastore.Get(id)

	if value == "" {
		_, ts, err := handler.slack.PostMessage("general", "Hello", messageParams)
		if err != nil {
			message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
		}

		err = handler.datastore.Set(id, ts)
		if err != nil {
			message := fmt.Sprintf("failed to set id: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
		}
	} else {
		messageParams.ThreadTimestamp = value
		_, _, err := handler.slack.PostMessage("general", "Hello", messageParams)
		if err != nil {
			message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusCreated)
}
