FROM golang:1.20.0-alpine

WORKDIR /app/go

COPY go.mod /app/go/
COPY go.sum /app/go/

RUN go mod download
RUN go mod tidy

COPY --chmod=765 . /app/go/
RUN chmod +x /app/go/

EXPOSE 4000

RUN ./build.sh

CMD ["./bin/go-diego-bot-linux"]
