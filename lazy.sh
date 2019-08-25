#!/bin/bash -e

REGISTRY_URL=${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com

aws configure set default.region ${AWS_REGION}

$(aws ecr get-login --no-include-email)

docker tag ${IMAGE_NAME} ${REGISTRY_URL}/${IMAGE_NAME}:latest

docker push ${REGISTRY_URL}/${IMAGE_NAME}:latest
