set -ex

docker-compose build
docker-compose -f ./web/docker-compose.yml build

docker-compose up -d
docker-compose -f ./web/docker-compose.yml up -d

docker-compose logs -f
docker-compose -f ./web/docker-compose.yml logs -f
