# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Build the Go app
RUN go build -o server ./cmd/app

# Expose the server port
EXPOSE 8080

# Run the server when the container launches
CMD ["./server"]