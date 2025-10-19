# Start with the official Go image as a base
FROM golang:1.22-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod ./
COPY go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 is used to disable CGO, which creates a statically linked binary
# -o froztypoint specifies the output binary name
# ./cmd/api is the path to your main packagecd 
RUN CGO_ENABLED=0 GOOS=linux go build -o froztypoint ./cmd/api

# Expose the port your application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./froztypoint"]
