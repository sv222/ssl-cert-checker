# Use an official Go runtime as a parent image
FROM golang:1.20.1-alpine3.17 as builder
# Set the working directory to /app
WORKDIR /app
# Copy the current directory contents into the container at /app
COPY . /app
# Build the main.go file
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Create the final container from prebuilt builder
FROM alpine:3.17
# Copy binary from builer to result directory
COPY --from=builder /app .
# Run the executable with the passed argument as the address file
ENTRYPOINT ["./app"]