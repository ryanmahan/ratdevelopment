#!/bin/sh
export GOPATH=/go
(go run ratdevelopment-backend --cassandra_ips 127.0.0.1) &
