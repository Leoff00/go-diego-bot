FROM golang:1.20.0-alpine

LABEL maintainer="leoff00"

RUN apk update && apk add --no-cache bash

WORKDIR /go/app

COPY . /go/app/

RUN go mod download && go mod tidy

COPY ./.env /go/app/

EXPOSE 4000

RUN go build -o /go/app/diegobot

CMD ["/go/app/diegobot"]
