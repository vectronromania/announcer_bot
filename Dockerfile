FROM golang:1.8

ADD . "~/go/src/github.com/cezarmathe/announcer-bot"

RUN go install github.com/cezarmathe/announcer-bot

ENTRYPOINT "~/go/bin/announcer-bot"
