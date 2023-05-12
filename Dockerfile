# Use an official Golang runtime as a parent image
FROM golang:latest

WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Initialize a new module
RUN go mod init goproject

# Install the GO dependencies
RUN go get -u fmt encoding/json log net/http

# Build the Go app
RUN go build -o binary .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./binary"]