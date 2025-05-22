FROM golang:1.24-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build both services
RUN go build -o /app/bin/api ./cmd/api
RUN go build -o /app/bin/worker ./cmd/worker

# Default fallback (can be overridden in docker-compose)
CMD ["/app/bin/api"]
