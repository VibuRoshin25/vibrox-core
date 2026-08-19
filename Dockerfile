FROM golang:1.26.6-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o main ./cmd

# Use a minimal image for running
FROM alpine:3.23

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/main .

RUN addgroup -S vibrox && adduser -S -G vibrox vibrox
USER vibrox

# Expose port  
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./main"]
