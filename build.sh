#!/bin/bash

set -ex
IMAGE_NAME="jincort/backend-cert2addr"
TAG="${1}"
docker run --rm -v "$PWD":/go/src/github.com/JincorTech/backend-cert2addr -w /go/src/github.com/JincorTech/backend-cert2addr golang:1.9.2 go test -v ./...
docker run --rm -v "$PWD":/go/src/github.com/JincorTech/backend-cert2addr -w /go/src/github.com/JincorTech/backend-cert2addr golang:1.9.2 go build -v
docker build -t ${IMAGE_NAME}:${TAG} .
docker push ${IMAGE_NAME}:${TAG}
