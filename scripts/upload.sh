#!/bin/bash
# Takes in the directory containing the datadumps
# Runs the mouse_upload program to load cassandra with the datadumps
# Resolve the directory of this script
SCRIPT_DIR=$(dirname $(readlink -f "$0"))
go run "$SCRIPT_DIR/mouse_upload/main.go" $1