FROM golang
WORKDIR /go/src/github.com/naoty/slack-thread-webhook
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/github.com/naoty/slack-thread-webhook/
RUN dep ensure

CMD ["slack-thread-webhook"]
