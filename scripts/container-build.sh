#!/usr/bin/env bash

TAG=${IMAGE_TAG:-latest}
CLI=${CONTAINER_CLI:-nerdctl}

${CLI} build -t ghcr.io/jdweeks/journey:${TAG} .