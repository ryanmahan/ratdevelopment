FROM golang:latest AS builder
RUN mkdir /ratdevelopment-latest
ADD . /ratdevelopment-latest/
WORKDIR /ratdevelopment-latest

RUN go get
RUN go build -o main .
CMD ["/app/main"]
