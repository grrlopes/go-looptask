# Use official Go image as the base image
FROM golang:1.23.1 AS builder

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum files first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# RUN go build -o /usr/local/bin/app ./src/server.go
# Build the Go binary statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o /usr/local/bin/app ./src/server.go

# Uncomment this line below CMD["app"] for debugging and comment the entire stage 2
# CMD ["app"]

# Stage 2: Create a lightweight final image without the source code
FROM alpine:latest

# Copy the Go binary from the builder stage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

RUN chmod +x /usr/local/bin/app

# Expose the port on which the app runs
EXPOSE 8080
EXPOSE 443

CMD ["/usr/local/bin/app"]
