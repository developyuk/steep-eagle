set -ex

docker-compose down
docker-compose -f ./web/docker-compose.yml down --remove-orphans
