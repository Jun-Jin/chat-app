#!/bin/sh

cmd="apply"

if [ $# -ne 0 ] && [ "$1" == "-d" ]; then
  cmd="delete"
fi

kubectl $cmd -f ./k8s/server/deployment.yaml
kubectl $cmd -f ./k8s/server/service.yaml
# kubectl $cmd -f ./k8s/server/ingress.yaml
