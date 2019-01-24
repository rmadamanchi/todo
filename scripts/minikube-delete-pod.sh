#!/usr/bin/env bash
kubectl delete pod $(kubectl get pods --selector=app=todo -o jsonpath='{.items[0].metadata.name}') || true
