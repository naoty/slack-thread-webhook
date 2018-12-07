package main

import (
	"os"
)

func main() {
	router := &router{}
	router.post("/hooks/(?P<id>\\d+)", handler)

	cli := cli{outStream: os.Stdout, errStream: os.Stderr, router: router}
	status := cli.run(os.Args)
	os.Exit(status)
}
