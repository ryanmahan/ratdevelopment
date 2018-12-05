#!/bin/sh
cd /home/ubuntu/go/src/ratdevelopment-backend
while ! cqlsh -e 'describe cluster' ; do
    sleep 1
done
cqlsh < init.cql
