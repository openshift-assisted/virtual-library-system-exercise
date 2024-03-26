# Use the official Go image as the base image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/webserver

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the config.yaml file from the builder stage
COPY --from=builder /app/config.yaml .

# Expose the port on which your application listens (if applicable)
EXPOSE 8080

# Run the application binary
CMD ["./main"]
