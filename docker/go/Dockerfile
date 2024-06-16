FROM golang:1.22.4 as build

RUN apt-get update

WORKDIR /go/src
COPY ./src /go/src

RUN go build -o go

RUN go install github.com/air-verse/air@latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/rubenv/sql-migrate/...@latest

RUN go install github.com/rakyll/statik@latest

CMD ["air"]
