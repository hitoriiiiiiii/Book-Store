# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (to leverage caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o bookstore ./cmd/main.go

# Stage 2: Minimal runtime image
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bookstore .

# Copy .env if you want environment variables inside container
COPY .env .env

# Expose port
EXPOSE 8080

# Command to run
CMD ["./bookstore"]
