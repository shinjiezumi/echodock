FROM golang:1.13.7 as build

RUN apt-get update && \
    apt-get install -y --no-install-recommends git

WORKDIR /go/src
COPY ./src /go/src

RUN go build -o go \
    && go get github.com/oxequa/realize \
#    && go get -u github.com/go-delve/delve/cmd/dlv \
#    && go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv
    && go get github.com/rubenv/sql-migrate/...