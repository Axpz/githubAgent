#!/bin/bash
IMAGE_NAME="hacknews"
VERSION="v0.0.1"
docker build --platform linux/arm64 -t ${IMAGE_NAME}:${VERSION} .
docker save -o ${IMAGE_NAME}:${VERSION}.linux-arm64.tar ${IMAGE_NAME}:${VERSION}
#kubectl delete Job hacknews-job
#kubectl apply -f hacknews-job.yaml
#kubectl logs -f -l job-name=hacknews-job