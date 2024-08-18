# Start from the official Go image
FROM golang:1.22-alpine as build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o myapp

# Use a minimal image to run the application
FROM alpine:3.18

# Copy the binary from the build stage
COPY --from=build /app/myapp /app/myapp

# Set the working directory
WORKDIR /app

# Expose the port on which the application runs
EXPOSE 3000

# Run the application
CMD ["/app/myapp"]
