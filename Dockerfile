# Use official Go image as a base
FROM golang:1.22.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o userManagementAPI main.go

# Use a runtime image with the required glibc version
FROM debian:bookworm-slim

# Set working directory inside the container
WORKDIR /

# Install glibc
RUN apt-get update && apt-get install -y libc6 && rm -rf /var/lib/apt/lists/*

# Copy the compiled binary from the builder
COPY --from=builder /app/userManagementAPI .

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["./userManagementAPI"]
