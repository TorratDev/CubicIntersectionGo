# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies (cache dependencies for faster subsequent builds)
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go web API binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create a lightweight runtime environment
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the Go API will run (e.g., 8080)
EXPOSE 8080

# Command to run the Go binary
CMD ["./main"]
