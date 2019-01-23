#!/usr/bin/env bash
set ex
export K8S_NAMESPACE=todo
kubectl get namespace ${K8S_NAMESPACE} >/dev/null || kubectl create namespace ${K8S_NAMESPACE}

kubectl delete pod todo

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
    image: todo:debug
    imagePullPolicy: Never
EOF

sleep 2s
kubetail todo