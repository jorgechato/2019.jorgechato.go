#!/bin/bash -e

REGISTRY_URL=${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com

aws configure set default.region ${AWS_REGION}

$(aws ecr get-login --no-include-email)

if [[ ${TRAVIS_BRANCH} == "master" ]] ; then
    PUSH_IMAGE=api-jorgechato-com:latest

    docker tag ${IMAGE_NAME} ${REGISTRY_URL}/${PUSH_IMAGE}
    docker push ${REGISTRY_URL}/${PUSH_IMAGE}
fi

docker tag ${IMAGE_NAME} ${REGISTRY_URL}/${IMAGE_NAME}
docker push ${REGISTRY_URL}/${IMAGE_NAME}
