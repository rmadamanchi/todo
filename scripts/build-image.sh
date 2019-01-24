#!/usr/bin/env bash
go mod vendor
docker build -t todo:latest -f ./build/package/Dockerfile .