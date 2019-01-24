#!/usr/bin/env bash
docker run -p 8080:8080 -p 10001:10001 --security-opt=seccomp:unconfined todo:debug