#!/bin/bash

APICLARITY_AGENT_IMAGE=apiclarity-bigip-agent

docker run -d \
    -v ${PWD}/../config.yaml:/config.yaml --env CONFIG_PATH=/config.yaml \
    -v ${PWD}/../apiclarity.crt:/apiclarity.crt \
    ${APICLARITY_AGENT_IMAGE}