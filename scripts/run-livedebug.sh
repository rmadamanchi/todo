#!/usr/bin/env bash
set -ex
docker run -p 8080:8080 -p 10001:10001 -v /goprojects/todo:/goprojects/todo --security-opt=seccomp:unconfined todo:livedebug