# Use a lightweight base image containing Go runtime
FROM golang:alpine AS builder

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PROTOC_VERSION=3.15.6

# Install protoc compiler
RUN apk add --no-cache curl unzip
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local && \
    rm -f protoc-${PROTOC_VERSION}-linux-x86_64.zip

# Install protoc Go plugin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Generate the protocol buffer files using Makefile
RUN make gen

# Load environment variables from .env file
ARG ENV_FILE
ENV ENV_FILE ${ENV_FILE}
RUN if [ -f "${ENV_FILE}" ]; then export $(cat ${ENV_FILE} | xargs); fi

# Build the Go binary
RUN go build -o app ./cmd/server

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app .

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./app"]
