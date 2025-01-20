#!/bin/bash
ctr -n=k8s.io images import /mnt/github.com/githubAgent/devices/hacknews/hacknews\:v0.0.1.linux-arm64.tar
kubectl delete Job hacknews-job
kubectl apply -f hacknews-job.yaml
kubectl apply -f hacknews-cronjob.yaml
kubectl get jobs