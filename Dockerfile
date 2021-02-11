FROM golang:latest AS builder

RUN apt update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /go/src/app
COPY go.mod .
RUN go mod download

COPY . . 
RUN go install
RUN go build



CMD ["./twitter-go"]
