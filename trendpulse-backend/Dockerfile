FROM golang:1.24-alpine

# Install required tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy and download dependencies first (for Docker cache)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Install gqlgen for codegen
RUN go install github.com/99designs/gqlgen@v0.17.73

# Generate GraphQL types (make sure gqlgen.yml + schema are copied)
RUN /go/bin/gqlgen generate

# Build both services
RUN mkdir -p /app/bin
RUN go build -o /app/bin/api ./cmd/api
RUN go build -o /app/bin/worker ./cmd/worker

# Default command (can be overridden)
CMD ["/app/bin/api"]
