package main

import (
	"bytes"
	"testing"
)

func TestHelp(t *testing.T) {
	cli := cli{outStream: bytes.NewBufferString(""), errStream: bytes.NewBufferString("")}
	status := cli.run([]string{"slack-thread-webhook", "--help"})
	if status != 0 {
		t.Errorf("expected: 0, got: %d\n", status)
	}
}

func TestVersion(t *testing.T) {
	cli := cli{outStream: bytes.NewBufferString(""), errStream: bytes.NewBufferString("")}
	status := cli.run([]string{"slack-thread-webhook", "--version"})
	if status != 0 {
		t.Errorf("expected: 0, got: %d\n", status)
	}
}
