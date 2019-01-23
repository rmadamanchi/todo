#!/usr/bin/env bash
set -ex
export K8S_NAMESPACE=todo
kubectl port-forward $(kubectl get pods --selector=app=todo -o jsonpath='{.items[0].metadata.name}') 8080:8080