#!/usr/bin/env bash
docker save todo:debug | (eval $(minikube docker-env) && docker load)