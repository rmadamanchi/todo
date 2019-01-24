#!/usr/bin/env bash
go mod vendor
docker build -t todo:debug -f ./build/package/Dockerfile.debug .