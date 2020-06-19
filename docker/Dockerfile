FROM golang:1.13.12-alpine3.12

RUN apk update && apk add bash

WORKDIR /go/src/main

COPY ./src/main/ .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT [ "/go/bin/main" ]