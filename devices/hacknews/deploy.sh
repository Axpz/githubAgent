#!/bin/bash
ctr -n=k8s.io images import /mnt/github.com/githubAgent/devices/hacknews/hacknews\:v0.0.1.linux-arm64.tar
kubectl delete Job -l app=hacknews
kubectl delete cronjob -l app=hacknews
kubectl apply -f hacknews-job.yaml
kubectl apply -f hacknews-cronjob.yaml
kubectl get job
kubectl get cronjob