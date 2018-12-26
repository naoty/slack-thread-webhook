# slack-thread-webhook
Slack webhook to post threaded messages

## Usage
slack-thread-webhook accepts requests at following endpoints. `<id>` is non-whitespace string.

* `POST /hooks/<id>` forwards requests to Slack. When a request with the same `<id>` has already sent, it will be posted as a theaded message.
* `PUT /hooks/<id>` updates a message posted with the same `<id>`.

## Setup

### Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### Docker
Docker images are published at [Docker Hub](https://hub.docker.com/r/naoty/slack-thread-webhook).
