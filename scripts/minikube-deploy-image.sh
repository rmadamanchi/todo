#!/usr/bin/env bash
set ex
docker save todo:latest | (eval $(minikube docker-env) && docker load)