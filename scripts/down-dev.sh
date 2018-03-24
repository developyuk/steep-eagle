set -ex

docker-compose down \
& docker-compose -f docker-compose-frontend.yml down --remove-orphans
