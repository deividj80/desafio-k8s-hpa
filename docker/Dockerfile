FROM golang:1.6.1-alpine

ENV APP_NAME main
ENV PORT 8080

COPY ./src/main/ /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}