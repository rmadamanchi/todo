#!/usr/bin/env bash
export K8S_NAMESPACE=todo
kubectl get namespace ${K8S_NAMESPACE} >/dev/null || kubectl create namespace ${K8S_NAMESPACE}

kubectl delete rs todo || true

cat <<EOF | kubectl create --namespace=${K8S_NAMESPACE} -f -
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: todo
  labels:
    app: todo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: todo
  template:
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
        - name: CGO_ENABLED
          value: "0"
        volumeMounts:
        - mountPath: "/goprojects/todo"
          name: goproject
      volumes:
      - name: goproject
        hostPath:
          path: "/goprojects/todo"
EOF

sleep 2s
./scripts/minikube-tail.sh