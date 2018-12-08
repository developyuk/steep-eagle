#!/usr/bin/env bash
set -x

docker-compose down
docker-compose build
docker-compose up -d
docker-compose logs --follow --tail 50
