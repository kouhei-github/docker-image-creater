FROM golang:1.16-alpine

RUN apk add --no-cache git

RUN go install github.com/spf13/cobra-cli@latest
RUN go get -u github.com/spf13/cobra@latest
RUN go get -u github.com/mitchellh/go-homedir@latest

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /var/www/application

COPY ./src /var/www/application

