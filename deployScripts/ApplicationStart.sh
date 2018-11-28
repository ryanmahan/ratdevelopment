#!/bin/sh
ls
cqlsh < init.cql
go get
go build
./ratdevelopment-backend
