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
    image: todo:livedebug
    imagePullPolicy: Never
    env:
    - name: GO111MODULE
      value: "on"
    - name: CGO_ENABLED
      value: "0"
    volumeMounts:
    - mountPath: "/goprojects"
      name: goprojects
  volumes:
  - name: goprojects
    hostPath:
      path: "/goprojects"
EOF

sleep 2s
kubetail todo