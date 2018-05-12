#!/usr/bin/env bash
set -x

docker-compose stop
docker stop emq220 && docker rm emq220
