set -ex

sudo docker-compose -f docker-compose-prod.yml stop
sudo docker-compose -f ./web/docker-compose-prod.yml stop
