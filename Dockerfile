# Use the official Golang image as a parent image
FROM golang:1.19-alpine3.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download Go dependencies
RUN go mod download

# Copy the source code for the app
COPY apps/weather ./apps/weather
COPY libs/weather-lib ./libs/weather-lib

# Build the app
RUN go build -o /app/weather ./apps/weather

# Use a small image for the final result
FROM alpine:3.17

# Copy the built executable from the builder stage
COPY --from=builder /app/weather /app/weather

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the app when the container starts
CMD ["/app/weather"]
