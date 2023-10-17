# Stage 1: Build the Go application
FROM golang:1.21.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go script and any necessary files
COPY . .

# Build the Go application
RUN go build -o item-store

# Stage 2: Create a smaller final image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/item-store .

# Expose the port the application will run on
EXPOSE 8080

# Define the command to run your application
CMD ["./item-store"]
