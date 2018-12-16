FROM golang AS builder
WORKDIR /go/src/github.com/naoty/slack-thread-webhook
RUN go get -u github.com/golang/dep/cmd/dep && go get -u github.com/pilu/fresh

ADD Gopkg.toml Gopkg.lock /go/src/github.com/naoty/slack-thread-webhook/
RUN dep ensure -vendor-only

ADD . /go/src/github.com/naoty/slack-thread-webhook/
RUN go install

FROM gcr.io/distroless/base
COPY --from=builder /go/bin/slack-thread-webhook /
CMD ["/slack-thread-webhook"]
