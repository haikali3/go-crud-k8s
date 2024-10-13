# Start with Golang base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Install dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the application
RUN go build -o /go-crud-k8s

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["/go-crud-k8s"]
