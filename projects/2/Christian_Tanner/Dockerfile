# Use the official Golang image to build the program
FROM golang:1.20-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go source code into the container
COPY hello.go .

# Build the Go application
RUN go build -o hello .

# Start a new image for the final container (minimal image)
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder image
COPY --from=builder /app/hello .

# Make the binary executable
RUN chmod +x hello

# Set the entry point to run the program
ENTRYPOINT ["./hello"]
