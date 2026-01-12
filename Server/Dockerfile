## DockerFile: For "Go" Application ##

## Stage-1: Build the "Go" application

# Start from the latest "golang" base image (alpine)
FROM golang:1.23.0-alpine3.20 AS build

# Set the 'app' as a 'CWD' inside the 'container'
WORKDIR /app

# Copy the 'go.mod' and 'go.sum' files first to 'leverage' the 'Docker' layer 'caching'
COPY go.mod go.sum ./

# Download all the 'dependencies' using the 'tidy' command
RUN go mod tidy

# Copy the rest of the "Go" application code
COPY . .

## Stage-2: Run the "Go" application

# Build the "Go" app
RUN go build -o main .

# Expose the port "3000" to the 'outside' world
EXPOSE 3000

# Run the 'binary' program
CMD ["/app/main"]