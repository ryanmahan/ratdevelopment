#!/bin/sh
killall go
sudo rm -rf /go
sudo service cassandra restart
