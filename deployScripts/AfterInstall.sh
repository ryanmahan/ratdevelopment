#!/bin/sh
ls
while ! cqlsh -e 'describe cluster' ; do
  sleep 1
done
cqlsh < ../init.cql
go get ../
go build -o ../ratdevelopment-backend ../
