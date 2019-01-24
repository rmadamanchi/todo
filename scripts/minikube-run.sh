#!/usr/bin/env bash
export K8S_NAMESPACE=todo
kubectl get namespace ${K8S_NAMESPACE} >/dev/null || kubectl create namespace ${K8S_NAMESPACE}

kubectl delete pod todo || true

cat <<EOF | kubectl create --namespace=${K8S_NAMESPACE} -f -
apiVersion: v1
kind: Pod
metadata:
  name: todo
  labels:
    app: todo
spec:
  containers:
  - name: todo
    image: todo:latest
    imagePullPolicy: Never
EOF

sleep 2s
./scripts/minikube-tail.sh