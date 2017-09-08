#!/bin/sh

# Exit on any error
set -e

SERVICE_NAME=$1
IMAGE_NAME=gcr.io/firsttry-178217/try-go-kit/$SERVICE_NAME
IMAGE_NAME_WITH_TAG=$IMAGE_NAME:$CIRCLE_SHA1
IMAGE_NAME_WITH_LATEST_TAG=$IMAGE_NAME:latest

cd $SERVICE_NAME && docker build -f Dockerfile -t $IMAGE_NAME_WITH_TAG .
${HOME}/google-cloud-sdk/bin/gcloud docker -- push $IMAGE_NAME_WITH_TAG
${HOME}/google-cloud-sdk/bin/gcloud container images add-tag $IMAGE_NAME_WITH_TAG $IMAGE_NAME_WITH_LATEST_TAG
${HOME}/google-cloud-sdk/bin/kubectl set image deployment $SERVICE_NAME $SERVICE_NAME=$IMAGE_NAME_WITH_TAG
