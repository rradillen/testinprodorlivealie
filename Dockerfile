# Start from a GoLang base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o testinprodorlivealie

# Expose a port (if your application listens on a specific port)
EXPOSE 8999

# Set the command to run the binary when the container starts
CMD ["./testinprodorlivealie"]
