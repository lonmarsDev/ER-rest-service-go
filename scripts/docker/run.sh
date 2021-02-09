#!/bin/bash
echo run th build
go mod download
docker build . -t go-er-service -f scripts/docker/Dockerfile
echo run docker
docker run -p 8080:8080 go-er-service