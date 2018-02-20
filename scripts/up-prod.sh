#!/bin/sh
set -ex

# docker-compose -f docker-compose-build.yml build web-main && \
# docker-compose -f docker-compose-build.yml up web-main && \

sudo docker-compose -f docker-compose-build.yml build db-main api-main db-admin web && \
sudo docker-compose -f docker-compose-build.yml up db-main api-main  db-admin web
