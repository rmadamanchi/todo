#!/usr/bin/env bash
export K8S_NAMESPACE=todo
kubectl port-forward $(kubectl get pods --selector=app=todo -o jsonpath='{.items[0].metadata.name}') 10001:10001
