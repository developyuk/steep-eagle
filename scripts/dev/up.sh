#!/usr/bin/env bash
set -x

docker stop emq220
docker rm emq220
docker run -tid --name emq220 -p 1883:1883 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqttd-docker-v2.2.0

docker-compose build
echo "waiting mqtt server ready to connect"
#sleep 1m
sleep 22
docker-compose up -d
docker-compose logs -f
