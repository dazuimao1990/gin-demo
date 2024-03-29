#!/bin/bash

BUILD_CMD="go build -o gin-demo"

docker run -ti --rm -v $(PWD):/code -w /code -e GOOS=linux golang:1.19 ${BUILD_CMD}