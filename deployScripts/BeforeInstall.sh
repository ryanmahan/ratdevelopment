#!/bin/sh
killall go
rm -rf /go > /goresult.txt
service cassandra restart
