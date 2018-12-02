VERSION ?= $$(git describe --tags)

install:
	go install -ldflags "-X main.Version=$(VERSION)"
