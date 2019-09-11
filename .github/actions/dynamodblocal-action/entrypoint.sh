#!/bin/sh

docker_run="docker run"

docker_run="$docker_run -d -p $INPUT_HOST_PORT:$INPUT_CONTAINER_PORT amazon/dynamodb-local"

sh -c "$docker_run"
