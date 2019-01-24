#!/usr/bin/env bash
docker save todo:livedebug | (eval $(minikube docker-env) && docker load)