#!/usr/bin/env bash
docker save todo:latest | (eval $(minikube docker-env) && docker load)