#!/bin/sh
while ! cqlsh -e 'describe cluster' ; do
  sleep 1
done
killall go
rm -rf /go > /goresult.txt
sudo service cassandra restart
