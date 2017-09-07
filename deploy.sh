#!/bin/bash

# Exit on any error
set -e

SERVICE_NAME=$1
cd $SERVICE_NAME && docker build -f Dockerfile -t gcr.io/firsttry-178217/try-go-kit/$SERVICE_NAME:$CIRCLE_SHA1 .
${HOME}/google-cloud-sdk/bin/gcloud docker -- push gcr.io/firsttry-178217/try-go-kit/$SERVICE_NAME:$CIRCLE_SHA1
${HOME}/google-cloud-sdk/bin/gcloud container images add-tag gcr.io/firsttry-178217/try-go-kit/$SERVICE_NAME:$CIRCLE_SHA1 gcr.io/firsttry-178217/try-go-kit/$SERVICE_NAME:latest
${HOME}/google-cloud-sdk/bin/kubectl set image deployment $SERVICE_NAME $SERVICE_NAME=gcr.io/firsttry-178217/try-go-kit/$SERVICE_NAME:$CIRCLE_SHA1
