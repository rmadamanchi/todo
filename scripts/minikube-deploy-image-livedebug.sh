#!/usr/bin/env bash
set ex
docker save todo:livedebug | (eval $(minikube docker-env) && docker load)