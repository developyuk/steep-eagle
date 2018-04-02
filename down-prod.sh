set -ex

docker-compose -f docker-compose-prod.yml stop
docker-compose -f ./web/docker-compose-prod.yml stop
