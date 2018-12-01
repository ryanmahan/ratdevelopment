#!/bin/sh
export GOPATH=/go
cd /go/src/ratdevelopment-backend
echo "=== GET STEP START ===" > /home/ubuntu/gobuildlogs.txt 2>&1
go get -v >>/home/ubuntu/gobuildlogs.txt 2>&1
echo "=== GET STEP DONE ===" >> /home/ubuntu/gobuildlogs.txt 2>&1
echo "=== BUILD STEP START ===" >> /home/ubuntu/gobuildlogs.txt 2>&1
go build >>/home/ubuntu/gobuildlogs.txt 2>&1
echo "=== BUILD STEP DONE ===" >> /home/ubuntu/gobuildlogs.txt 2>&1
echo "=== INSTALL STEP START ===" >> /home/ubuntu/gobuildlogs.txt 2>&1
go install >>/home/ubuntu/gobuildlogs.txt 2>&1
echo "=== INSTALL STEP DONE ===" >> /home/ubuntu/gobuildlogs.txt 2>&1
