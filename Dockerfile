# Start by building the Go application
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Now create a smaller container for running the app
FROM gcr.io/distroless/base-debian10

# Copy the compiled binary from the build stage
COPY --from=build /app/main /app/main

# Expose the port the app will run on
EXPOSE 4000

# Command to run the binary
CMD ["/app/main"]
