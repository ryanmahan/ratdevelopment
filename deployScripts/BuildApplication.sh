#!/bin/sh
export GOPATH=/go
cd /go/src/ratdevelopment-backend
sudo rm ~/gobuildlogs.txt
touch ~/gobuildlogs.txt
go get 2>>~/gobuildlogs.txt
go build 2>>~/gobuildlogs.txt
go install 2>>~/gobuildlogs.txt
