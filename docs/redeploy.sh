#!/bin/bash
set -ex

kubectl delete deploy github-agent-server
ctr -n=k8s.io images import /mnt/github.com/githubAgent/github-agent-server\:v0.0.1.linux-arm64.tar
kubectl apply -f /mnt/github.com/githubAgent/docs/github-agent-server.yaml

sleep 3
kubectl logs -f -l app=github-agent-server