# Use an official Golang runtime as a parent image
FROM golang:1.21.4

# Set the working directory inside the container
WORKDIR /app

COPY . .
# Download Go module dependencies and build the application
RUN go mod tidy
RUN go build -o main

# Expose the port the application will run on
EXPOSE 8080

# Run the application
CMD ["./main"]
