#!/bin/sh

# Exits immediately if a command exits with a non-zero status
set -e

# Run 'docker-compose up' for source database deployment
docker-compose -f docker/db/docker-compose.yml up -d

# Build the Go application
go build -o gin_api_with_auth src/main.go

# Run the application
./gin_api_with_auth