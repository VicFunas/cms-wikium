# Use the official Golang image to create a build artifact.
FROM golang:1.24-bullseye AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies.
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app for a Linux environment
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/server ./cmd/server

# Start a new stage from scratch for a smaller final image
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/server .

# Expose port 50051 to the network
EXPOSE 50051

# Command to run the executable
CMD ["./server"]
