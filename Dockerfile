# Use the official Go image as the base image
FROM golang:1.21 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files to the working directory
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o /app/app

# Expose the port your Go application is listening on
EXPOSE 8080

# Command to run the Go application
CMD ["./app"]