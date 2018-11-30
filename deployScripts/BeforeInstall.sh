#!/bin/sh
killall go
sudo rmdir /go
sudo service cassandra restart
