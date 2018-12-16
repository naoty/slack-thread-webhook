package wrapper

import (
	"fmt"
	"strings"
	"net/http"
)

// WithJSON returns a handler to check that the content type is application/json.
func WithJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			h.ServeHTTP(w, req)
			return
		}

		message := fmt.Sprintf("Content-Type must be application/json: %s", contentType)
		http.Error(w, message, http.StatusBadRequest)
	})
}
