FROM golang:1.20-buster

WORKDIR /usr/app
RUN go install github.com/cosmtrek/air@latest

COPY src/go.mod .
RUN go clean --modcache

RUN go mod tidy && go mod vendor