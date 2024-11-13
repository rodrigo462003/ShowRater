FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go install github.com/air-verse/air@latest

# Expose port
EXPOSE 8080

# Command to run the application with Air for auto-reload
CMD ["air"]

