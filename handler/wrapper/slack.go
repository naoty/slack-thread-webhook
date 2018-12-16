package wrapper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
)

// PostMessageParametersKey is a constant to fetch parameters for Slack from context.
const PostMessageParametersKey contextKey = "PostMessageParametersKey"

// WithSlack is a wrapper to parse request body to arguments for Slack API and
// store them in context.
func WithSlack(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Body == nil {
			http.Error(w, "request body has nothing", http.StatusBadRequest)
			return
		}

		var parameters slack.PostMessageParameters
		err := json.NewDecoder(req.Body).Decode(&parameters)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), PostMessageParametersKey, parameters)
		newRequest := req.WithContext(ctx)
		h.ServeHTTP(w, newRequest)
	})
}
