package handler

import (
	"fmt"
	"net/http"

	"github.com/naoty/slack-thread-webhook/datastore"
	"github.com/naoty/slack-thread-webhook/handler/wrapper"
	"github.com/nlopes/slack"
)

// Put is a handler to update messages.
type Put struct {
	Channel   string
	Datastore datastore.Client
	Slack     *slack.Client
}

func (handler Put) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requestParams := req.Context().Value(wrapper.ParametersKey).(map[string]string)
	id := requestParams["id"]
	timestamp, _ := handler.Datastore.Get(id)

	if timestamp == "" {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	messageParams := req.Context().Value(wrapper.PostMessageParametersKey).(slack.PostMessageParameters)
	_, _, _, err := handler.Slack.SendMessage(
		handler.Channel,
		slack.MsgOptionUpdate(timestamp),
		slack.MsgOptionAttachments(messageParams.Attachments...),
		slack.MsgOptionPostMessageParameters(messageParams),
	)

	if err != nil {
		message := fmt.Sprintf("failed to post a message: %v\n", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
