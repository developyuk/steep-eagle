#!/bin/sh
set -ex

# docker-compose -f docker-compose-build.yml build web-main && \
# docker-compose -f docker-compose-build.yml up web-main && \

sudo docker-compose -f docker-compose-build.yml build web api-main db-main db-admin && \
sudo docker-compose -f docker-compose-build.yml up web api-main db-main db-admin
