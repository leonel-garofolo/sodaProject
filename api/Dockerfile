# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.16-alpine

# Add Maintainer Info
LABEL maintainer="Leonel Garofolo <garofolo.leonel@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

ARG COMMIT
ENV COMMIT ${COMMIT}

# Expose port 3000   to the outside world
EXPOSE 3000

# Run the executable
CMD ["/app/main"]