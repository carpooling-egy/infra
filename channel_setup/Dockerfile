# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o setup-nats main.go

# Stage 2: Run only the binary in a small container
FROM alpine:latest

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/setup-nats .

# Set binary as entrypoint
ENTRYPOINT ["./setup-nats"]