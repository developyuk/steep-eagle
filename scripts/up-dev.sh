set -ex

docker-compose build
docker-compose up -d

docker-compose -f docker-compose-frontend.yml build
docker-compose -f docker-compose-frontend.yml up -d
