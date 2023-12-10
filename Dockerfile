# Build stage 
FROM golang:1.21-alpine AS build-env

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . $(GOPATH)/src/github.com/Tiburso/GoManager
WORKDIR $(GOPATH)/src/github.com/Tiburso/GoManager

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

FROM alpine:latest

# Expose port 8080 to the outside world
EXPOSE 8080

# Copy the Pre-built binary file from the previous stage
COPY --from=build-env /go/bin/GoManager /go/bin/GoManager

# Setup a new group and user
RUN addgroup -S -g 1000 gomanager && \
    adduser -S -H -D -u 1000 -G gomanager gomanager && \
    chown -R gomanager:gomanager /go/bin/GoManager && \
    chmod +x /go/bin/GoManager && \
    echo "gomanager:*" | chpasswd -e

# Use an unprivileged user.
USER gomanager

# Command to run the executable
ENTRYPOINT ["/go/bin/GoManager"]


