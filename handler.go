package main

import (
	"fmt"
	"net/http"
)

var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	params := req.Context().Value(paramsKey).(map[string]string)
	for key, value := range params {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
})
