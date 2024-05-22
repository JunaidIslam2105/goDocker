# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Unknown27342 <unknown27342@example.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD ["./main"]