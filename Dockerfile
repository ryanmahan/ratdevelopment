FROM golang:latest AS builder
RUN mkdir -p /go/src/ratdevelopment-backend
ADD . /go/src/ratdevelopment-backend/
WORKDIR /go/src/ratdevelopment-backend

RUN go get
RUN go build -o main .
CMD ["./main"]
