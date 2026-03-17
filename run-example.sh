#!/bin/bash

# Load environment variables
set -a
source .env
set +a

# Run the example
go run example/main.go
