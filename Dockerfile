# Start from Go official image
FROM golang:1.25-alpine

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the application
RUN go build -o main .

# Run the executable
CMD ["./main"]
