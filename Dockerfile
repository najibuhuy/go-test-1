# Build Stage
FROM golang:1.22 AS builder
WORKDIR /app

# Copy dependency files and install modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Runtime Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the built binary and .env file
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env 

# Ensure correct permissions
RUN chmod +x /app/main

# Expose port
EXPOSE 9001

# Run the binary
CMD ["/app/main"]
