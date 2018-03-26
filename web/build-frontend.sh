set -ex


docker-compose -f ./docker-compose.yml build

docker-compose -f ./docker-compose.yml exec tutor npm run build
docker-compose -f ./docker-compose.yml exec operation npm run build

