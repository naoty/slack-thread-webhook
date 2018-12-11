package handler

import (
	"fmt"
	"net/http"

	"github.com/naoty/slack-thread-webhook/datastore"

	"github.com/nlopes/slack"
)

type contextKey int

const (
	paramsKey contextKey = iota
)

// Post is a handler to post messages.
type Post struct {
	Datastore datastore.Client
	Slack     *slack.Client
}

func (handler Post) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	messageParams := slack.NewPostMessageParameters()
	requestParams := req.Context().Value(paramsKey).(map[string]string)
	id := requestParams["id"]
	value, _ := handler.Datastore.Get(id)

	if value == "" {
		_, ts, err := handler.Slack.PostMessage("general", "Hello", messageParams)
		if err != nil {
			message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
		}

		err = handler.Datastore.Set(id, ts)
		if err != nil {
			message := fmt.Sprintf("failed to set id: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
		}
	} else {
		messageParams.ThreadTimestamp = value
		_, _, err := handler.Slack.PostMessage("general", "Hello", messageParams)
		if err != nil {
			message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusCreated)
}

