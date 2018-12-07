package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	router := &router{}
	router.post("/hooks/(?P<id>\\d+)", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.Context().Value(paramsKey).(map[string]string)
		for key, value := range params {
			fmt.Printf("key: %v, value: %v\n", key, value)
		}
	}))

	cli := cli{outStream: os.Stdout, errStream: os.Stderr, router: router}
	status := cli.run(os.Args)
	os.Exit(status)
}
