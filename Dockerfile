# Start from the latest Golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Your Name <your.email@domain.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

