FROM golang:1.17.7-bullseye

WORKDIR /go/src/app

RUN go install golang.org/x/tools/gopls@latest
