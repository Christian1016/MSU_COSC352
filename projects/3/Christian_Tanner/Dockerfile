# Use the official Golang image to build the program
FROM golang:1.20-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go source code into the container
COPY extract_tables.go .

# Download the required Go packages
RUN go get golang.org/x/net/html

# Build the Go application
RUN go build -o extract_tables .

# Start a new image for the final container (minimal image)
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder image
COPY --from=builder /app/extract_tables .

# Make the binary executable
RUN chmod +x extract_tables

# Set the entry point to run the program
ENTRYPOINT ["./extract_tables"]
