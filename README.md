# Gin API with auth middleware

## Abstract

This repository contains a Go-based API developed using the Gin web framework. 
The API features robust authentication and authorization mechanisms, allowing secure access to various endpoints. 
Users can create accounts, log in, and interact with wine-related data while adhering to specific access rights.


## Middleware

The API utilizes the following middleware functions:

* Authenticate: Validate users by querying DB with basic auth headers contained in request. Used before 'login' handler in user context.

* Authorize: Manages user access control based on JWT token claims. Used before handlers in wine context.
    
* LogTokenActivity: Logs user activity by creating token activity records for auditing purposes. Used after handlers in wine context.


## User deletion goroutine

A goroutine will be launched on startup which periodically (every 10 seconds) checks for users in the database who have been designated for removal. 
That is, users which has the 'for_deletion' flag set to 'true.' 
If any such users are found, they are deleted from the database. 

## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)

## Startup

The script "up" provisions resources and starts our application by executing the following:
```
1. docker-compose -f docker/db/docker-compose.yml up -d
2. go build -o gin_api_with_auth src/main.go
3. ./gin_api_with_auth
```

## Shutdown

The script "down" deletes our dev and test databases by executing the following:
```
1. docker-compose -f db/docker-compose.yml down
```

## Postman Collection

The repository includes a Postman collection in the 'postman' directory.