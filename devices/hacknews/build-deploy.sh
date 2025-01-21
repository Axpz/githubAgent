#!/bin/bash
docker build -t hacknews:v0.0.1 .
kubectl delete Job hacknews-job
kubectl apply -f hacknews-job.yaml
kubectl logs -f -l job-name=hacknews-job