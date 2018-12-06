package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	router := &router{}
	router.post("/hooks/(?P<id>\\d+)", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v\n", r)
	}))

	cli := cli{outStream: os.Stdout, errStream: os.Stderr, router: router}
	status := cli.run(os.Args)
	os.Exit(status)
}
