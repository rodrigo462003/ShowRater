FROM golang:latest

WORKDIR /app

RUN go install github.com/air-verse/air@latest

ENTRYPOINT ["air"]

