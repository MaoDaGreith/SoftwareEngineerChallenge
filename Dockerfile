# Start from the official Golang image
FROM golang:1.21-alpine as builder

# Set working directory
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Copy source code
COPY . .

# Build the application
RUN cd cmd/server && go build -o /app/server

# Use a minimal image for running
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy binary and web files
COPY --from=builder /app/server ./server
COPY --from=builder /app/web ./web

# Create non-root user for security
RUN adduser -D -s /bin/sh appuser
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1

# Run the application
CMD ["./server"] 