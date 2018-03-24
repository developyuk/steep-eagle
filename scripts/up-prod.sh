#!/bin/sh
set -ex

sudo docker-compose -f docker-compose-prod.yml build
sudo docker-compose -f docker-compose-prod.yml up -d

sudo docker-compose -f docker-compose-prod-frontend.yml build
sudo docker-compose -f docker-compose-prod-frontend.yml up -d


sudo docker-compose logs -f \
& sudo docker-compose -f docker-compose-frontend.yml logs -f

tail -f /dev/null
