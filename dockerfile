# Use an official Golang image with version 1.23
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to access the application
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
