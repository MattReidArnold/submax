# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Matt Arnold <matt.arnold@idmission.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service ./internal/cmd/service

######## Start a new stage from scratch #######
# FROM golang:latest  

FROM alpine:latest  

RUN apk update && apk --no-cache add ca-certificates curl bash redis

WORKDIR /root

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/service .

# Expose port 8081 to the outside
EXPOSE 8081

# Command to run the executable
CMD ["./service"] 
