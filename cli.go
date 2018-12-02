package main

import (
	"fmt"
	"io"
)

type cli struct {
	outStream, errStream io.Writer
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

func (cli cli) Run(args []string) int {
	if len(args) == 1 {
		fmt.Fprintln(cli.outStream, "TODO: implement")
		return exitCodeOK
	}

	switch args[1] {
	case "-h", "--help":
		fmt.Fprintf(cli.outStream, "%v", helpMessage)
		return exitCodeOK
	case "-v", "--version":
		fmt.Fprintln(cli.outStream, Version)
		return exitCodeOK
	}

	return exitCodeOK
}
