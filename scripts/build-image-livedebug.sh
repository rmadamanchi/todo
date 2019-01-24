#!/usr/bin/env bash
go mod vendor
docker build -t todo:livedebug -f ./build/package/Dockerfile.livedebug .