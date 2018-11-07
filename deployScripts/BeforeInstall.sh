#!/bin/sh
service cassandra stop
service cassandra start
cqlsh init.cql
