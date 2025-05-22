# Use the official Golang image (version 1.23) as the build stage
FROM golang:1.23 as build

# Set the working directory inside the container to /app
WORKDIR /app


# Copy all files from the current host directory into the container's /app directory
COPY . .

# Build the Go application from main.go and output the binary as 'main'
RUN go build -o main main.go

# List the contents of the /app directory (for debugging or verification purposes)
RUN ls -l /app

# Set the default command to run the compiled binary when the container starts
CMD ["./main"]
