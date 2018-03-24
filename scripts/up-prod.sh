#!/bin/sh
set -ex

sudo docker-compose -f docker-compose-prod.yml build
sudo docker-compose -f docker-compose-prod.yml up -d

sudo docker-compose -f docker-compose-prod-frontend.yml build
sudo docker-compose -f docker-compose-prod-frontend.yml up -d
