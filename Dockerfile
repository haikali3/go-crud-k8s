# First stage: build the Go application
FROM golang:1.22 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code to the container
COPY . .


# Second stage: create a minimal image to run the application
# Using a smaller image like Alpine for the runtime environment
FROM alpine:latest  

# Expose port 8000 for the application
EXPOSE 8000

# Run the Go app
CMD ["/go-crud-k8s"]
