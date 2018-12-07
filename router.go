package main

import (
	"context"
	"net/http"
	"regexp"
)

type router struct {
	routes map[string]map[*regexp.Regexp]http.Handler
}

type contextKey int

const (
	paramsKey contextKey = iota
)

func (router *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	paths, ok := router.routes[req.Method]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var handler http.Handler
	for re, h := range paths {
		matched := re.FindStringSubmatch(req.URL.Path)
		params := make(map[string]string)

		for i, name := range re.SubexpNames() {
			if i > 0 {
				params[name] = matched[i]
			}
		}

		handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), paramsKey, params)
			req = req.WithContext(ctx)
			h.ServeHTTP(w, req)
		})
	}

	if handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler.ServeHTTP(w, req)
}

func (router *router) post(pattern string, handler http.Handler) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return
	}

	if router.routes == nil {
		router.routes = make(map[string]map[*regexp.Regexp]http.Handler)
	}

	_, ok := router.routes["POST"]
	if !ok {
		router.routes["POST"] = make(map[*regexp.Regexp]http.Handler)
	}

	router.routes["POST"][re] = handler
}
