#!/bin/sh
export GOPATH=/go
cd /go/src/ratdevelopment-backend
sudo rm /home/ubuntu/gobuildlogs.txt
touch /home/ubuntu/gobuildlogs.txt
go get 2>>/home/ubuntu/gobuildlogs.txt
go build 2>>/home/ubuntu/gobuildlogs.txt
go install 2>>/home/ubuntu/gobuildlogs.txt
