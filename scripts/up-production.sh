set -ex

sudo docker-compose -f docker-compose-build.yml build web-main web-operation && \
sudo docker-compose -f docker-compose-build.yml up web-main web-operation && \

sudo docker-compose -f docker-compose-build.yml build web api-main db-main && \
sudo docker-compose -f docker-compose-build.yml up web api-main db-main
