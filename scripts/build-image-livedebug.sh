#!/usr/bin/env bash
set -ex
go mod vendor
docker build -t todo:livedebug -f ./build/package/Dockerfile.livedebug .