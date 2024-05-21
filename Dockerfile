# Use the official Golang image
FROM golang:1.16.3-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first
# This ensures that the go mod download step is cached if these files haven't changed
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o api .

# Expose the port on which the app will run
EXPOSE 8000

# Command to run the executable
CMD ["./api"]
