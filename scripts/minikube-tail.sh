#!/usr/bin/env bash
echo "Tailing Logs..."
kubectl logs --follow $(kubectl get pods --selector=app=todo -o jsonpath='{.items[0].metadata.name}')
