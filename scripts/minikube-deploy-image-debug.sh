#!/usr/bin/env bash
set ex
docker save todo:debug | (eval $(minikube docker-env) && docker load)