#!/bin/bash
# Migrate the schema to the Cassandra DB
# Resolve the directory of this script
SCRIPT_DIR=$(dirname $(readlink -f "$0"))
cqlsh --file "$SCRIPT_DIR/schema.cql"
