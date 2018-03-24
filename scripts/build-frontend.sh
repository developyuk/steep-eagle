set -ex


docker-compose -f docker-compose-frontend.yml build

docker-compose -f docker-compose-frontend.yml exec web-main npm run build
docker-compose -f docker-compose-frontend.yml exec web-operation npm run build

