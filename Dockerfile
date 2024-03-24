# Use the official Golang image as the base image
FROM golang:1.22.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o fixy-finance-api ./cmd/main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage to the final container
COPY --from=builder /app/fixy-finance-api .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./fixy-finance-api"]
