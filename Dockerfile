FROM golang:latest

WORKDIR /app

RUN apt-get update && \
    apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g tailwindcss

RUN go install github.com/air-verse/air@latest

ENTRYPOINT ["air"]

