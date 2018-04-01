set -ex


docker-compose -f ./web/docker-compose.yml build

docker-compose -f ./web/docker-compose.yml run tutor npm run build
docker-compose -f ./web/docker-compose.yml run operation npm run build

