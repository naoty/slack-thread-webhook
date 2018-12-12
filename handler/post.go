package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/naoty/slack-thread-webhook/datastore"

	"github.com/nlopes/slack"
)

type contextKey int

const (
	paramsKey contextKey = iota
)

// Post is a handler to post messages.
type Post struct {
	Channel string
	Datastore datastore.Client
	Slack     *slack.Client
}

func (handler Post) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		message := fmt.Sprintf("wrong Content-Type: %v", contentType)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	messageParams := slack.NewPostMessageParameters()
	requestParams := req.Context().Value(paramsKey).(map[string]string)
	id := requestParams["id"]
	value, _ := handler.Datastore.Get(id)

	if value == "" {
		_, ts, err := handler.Slack.PostMessage(handler.Channel, "Hello", messageParams)
		if err != nil {
			message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		err = handler.Datastore.Set(id, ts)
		if err != nil {
			message := fmt.Sprintf("failed to set id: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}
	} else {
		messageParams.ThreadTimestamp = value
		_, _, err := handler.Slack.PostMessage(handler.Channel, "Hello", messageParams)
		if err != nil {
			message := fmt.Sprintf("failed to post a message to slack: %v\n", err)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
