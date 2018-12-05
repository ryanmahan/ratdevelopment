#!/bin/sh
export GOPATH=/go
cd /go/src/ratdevelopment
echo "=== GET STEP START ===" &> /home/ubuntu/gobuildlogs.txt
go get -v &>> /home/ubuntu/gobuildlogs.txt
echo "=== GET STEP DONE ===" &>> /home/ubuntu/gobuildlogs.txt
echo "=== BUILD STEP START ===" &>> /home/ubuntu/gobuildlogs.txt
go build &>> /home/ubuntu/gobuildlogs.txt
echo "=== BUILD STEP DONE ===" &>> /home/ubuntu/gobuildlogs.txt
echo "=== INSTALL STEP START ===" &>> /home/ubuntu/gobuildlogs.txt
go install &>> /home/ubuntu/gobuildlogs.txt
echo "=== INSTALL STEP DONE ===" &>> /home/ubuntu/gobuildlogs.txt
