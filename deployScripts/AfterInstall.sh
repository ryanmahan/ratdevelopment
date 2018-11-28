#!/bin/sh
while ! cqlsh -e 'describe cluster' ; do
  sleep 1
done
export GOPATH=/go
cd /go/src/ratdevelopment-backend
cqlsh < init.cql
go get
go build
go install
