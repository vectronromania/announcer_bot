FROM golang:1.8

ADD . "~/go/src/github.com/cezarmathe/announcer_bot"

RUN go install github.com/cezarmathe/announcer_bot

ENTRYPOINT "~/go/bin/announcer_bot"
