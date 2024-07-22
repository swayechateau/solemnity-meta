# Use the official Golang image as a base image
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./app/. .

# Build the Go app
ENV CGO_ENABLED=0
RUN go build -o meta .

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/meta .
COPY --from=builder /app/public ./public


# Expose port 5050 to the outside world
EXPOSE 8080

USER 1000:1000

# Command to run the executable
CMD ["./meta"]
