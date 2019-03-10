#!/bin/bash

if [ ! -d "out" ]; then
    mkdir out
fi

env GOOS=linux GOARCH=amd64 GOARM=5 go build -v -o out/api .
docker build -t registry.heroku.com/docker-travis-heroku/web . -f build/Dockerfile