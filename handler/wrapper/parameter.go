package wrapper

import (
	"context"
	"net/http"
)

type contextKey string

// ParametersKey is a constant to fetch parameters from context.
const ParametersKey contextKey = "ParametersKey"

// WithParameters returns a wrapped handler with given parameters in context.
func WithParameters(h http.Handler, parameters map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), ParametersKey, parameters)
		newRequest := req.WithContext(ctx)
		h.ServeHTTP(w, newRequest)
	})
}
