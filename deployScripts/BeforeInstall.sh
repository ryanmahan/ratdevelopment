#!/bin/sh
killall go
rm -rf /go/src
rm -rf /go
sudo service cassandra restart
