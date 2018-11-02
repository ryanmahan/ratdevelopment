#!/bin/bash
# Migrate the schema to the Cassandra DB
cqlsh --file '/home/vagrant/go/schema.cql'
