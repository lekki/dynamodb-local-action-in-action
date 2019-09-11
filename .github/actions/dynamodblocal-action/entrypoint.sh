#!/bin/sh

docker_run="docker run"

docker_run="$docker_run -d -p 8000:8000 amazon/dynamodb-local"

sh -c "$docker_run"
