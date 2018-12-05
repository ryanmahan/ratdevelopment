#!/bin/sh
cd /home/ubuntu/go/src/ratdevelopment
while ! cqlsh -e 'describe cluster' ; do
    sleep 1
done
cqlsh < init.cql
