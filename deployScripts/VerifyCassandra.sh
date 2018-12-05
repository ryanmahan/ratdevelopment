#!/bin/sh
killall go
sudo rm -rf /home/ubuntu/go/src/ratdevelopment 2> /goresult.txt
sudo service cassandra restart
while ! cqlsh -e 'describe cluster' ; do
    sleep 1
done
