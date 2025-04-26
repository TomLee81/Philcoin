#!/usr/bin/env bash
set -e
IMAGE_NAME=philcoin-backend
TAG=${1:-latest}

GOOS=linux GOARCH=amd64 go build -o bin/server ./cmd/server

docker build -t ${IMAGE_NAME}:${TAG} .
docker push ${IMAGE_NAME}:${TAG}