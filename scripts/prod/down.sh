#!/usr/bin/env bash
set -x

sudo docker-compose -f docker-compose-prod.yml stop
sudo docker stop emq220
sudo docker rm emq220
