FROM golang
WORKDIR /go/src/github.com/naoty/slack-thread-webhook
RUN go get -u github.com/golang/dep/cmd/dep

ADD Gopkg.toml Gopkg.lock /go/src/github.com/naoty/slack-thread-webhook/
RUN dep ensure -vendor-only

ADD . /go/src/github.com/naoty/slack-thread-webhook/
RUN go install

CMD ["slack-thread-webhook"]
