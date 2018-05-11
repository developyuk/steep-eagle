#!/usr/bin/env bash
set -x

sudo docker stop emq220
sudo docker rm emq220
sudo docker run -tid --name emq220 -p 1883:1883 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqttd-docker-v2.2.0

sudo docker-compose -f docker-compose-prod.yml build
echo "waiting mqtt server ready to connect"
#sleep 1m
sleep 13
sudo docker-compose -f docker-compose-prod.yml up -d
sudo docker-compose -f docker-compose-prod.yml logs -f
