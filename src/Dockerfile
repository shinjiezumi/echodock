FROM golang:1.15.12-alpine3.12 as builder

RUN apk update \
    && apk add --no-cache git curl

WORKDIR /go/src/

COPY . .

RUN go mod download

# Build the command inside the container.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o api

# Use a Docker multi-stage build to create a lean production image.
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/api /api

ENV PORT=${PORT}

# Run the web service on container startup.
CMD ["/api"]