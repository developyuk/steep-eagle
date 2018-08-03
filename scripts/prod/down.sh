#!/usr/bin/env bash
set -x

docker-compose -f docker-compose-prod.yml stop
# docker stop emq220 && sudo docker rm emq220
