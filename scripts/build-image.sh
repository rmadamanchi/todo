#!/usr/bin/env bash
set -ex
go mod vendor
docker build -t todo:latest -f ./build/package/Dockerfile .