set -ex

docker-compose build
docker-compose -f docker-compose-frontend.yml build

docker-compose up -d
docker-compose -f docker-compose-frontend.yml up -d

docker-compose logs -f \
& docker-compose -f docker-compose-frontend.yml logs -f

tail -f /dev/null
