#!/usr/bin/env bash
set -ex
go mod vendor
docker build -t todo:debug -f ./build/package/Dockerfile.debug .