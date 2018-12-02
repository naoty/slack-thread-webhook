package main

import (
	"os"
)

func main() {
	cli := cli{outStream: os.Stdout, errStream: os.Stderr}
	status := cli.Run(os.Args)
	os.Exit(status)
}
