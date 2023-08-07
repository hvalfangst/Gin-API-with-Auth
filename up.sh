#!/bin/sh

# Exits immediately if a command exits with a non-zero status
set -e

# Run 'docker-compose up' for source database deployment
docker-compose -f docker/db/docker-compose.yml up -d

# Build the Go application
go build -o myapp src/main.go

# Run the application
./myapp