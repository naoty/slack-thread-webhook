package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/naoty/slack-thread-webhook/handler"
)

type cli struct {
	outStream, errStream io.Writer
	router               *handler.Router
}

const (
	exitCodeOK = 0
)

const helpMessage = `usage:
  slack-thread-webhook
  slack-thread-webhook (-h | --help)

options:
  -h, --help  show version.
`

// Version is the version of this application.
var Version string

func (cli cli) run(args []string) int {
	if len(args) > 1 {
		switch args[1] {
		case "-h", "--help":
			fmt.Fprintf(cli.outStream, "%v", helpMessage)
			return exitCodeOK
		case "-v", "--version":
			fmt.Fprintln(cli.outStream, Version)
			return exitCodeOK
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Fprintf(cli.outStream, "HTTP server started on :%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), cli.router); err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return exitCodeOK
	}

	return exitCodeOK
}
