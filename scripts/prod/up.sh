#!/usr/bin/env bash
set -x

sh ./scripts/prod/down.sh
# docker run -tid --name emq220 \
#  --volume /etc/letsencrypt:/etc/letsencrypt \
#  --env EMQ_DASHBOARD__LISTENER__HTTPS=18084 --env EMQ_DASHBOARD__LISTENER__HTTPS__ACCESS__1="allow all" \
#  --env EMQ_DASHBOARD__LISTENER__HTTPS__ACCEPTORS=2 --env EMQ_DASHBOARD__LISTENER__HTTPS__MAX_CLIENTS=512 \
#  --env EMQ_DASHBOARD__LISTENER__HTTPS__ACCESS__1="allow all" \
#  --env EMQ_DASHBOARD__LISTENER__HTTPS__KEYFILE="\/etc\/letsencrypt\/live\/mtutor.codingcamp.id\/privkey.pem" \
#  --env EMQ_DASHBOARD__LISTENER__HTTPS__CERTFILE="\/etc\/letsencrypt\/live\/mtutor.codingcamp.id\/fullchain.pem" \
#  --env EMQ_LISTENER__WSS__EXTERNAL__KEYFILE="\/etc\/letsencrypt\/live\/mtutor.codingcamp.id\/privkey.pem" \
#  --env EMQ_LISTENER__WSS__EXTERNAL__CERTFILE="\/etc\/letsencrypt\/live\/mtutor.codingcamp.id\/fullchain.pem" \
#  --env EMQ_LISTENER__SSL__EXTERNAL__KEYFILE="\/etc\/letsencrypt\/live\/mtutor.codingcamp.id\/privkey.pem" \
#  --env EMQ_LISTENER__SSL__EXTERNAL__CERTFILE="\/etc\/letsencrypt\/live\/mtutor.codingcamp.id\/fullchain.pem" \
#  --publish 1883:1883 --publish 8883:8883 \
#  --publish 8083:8083 --publish 8084:8084 \
#  --publish 18083:18083 --publish 18084:18084 \
#  emqttd-docker-v2.2.0
# echo "waiting mqtt server ready to connect"
# #sleep 1m
# sleep 13

docker-compose -f docker-compose-prod.yml build
docker-compose -f docker-compose-prod.yml up -d --scale api=5 --remove-orphans
docker-compose -f docker-compose-prod.yml logs --follow --tail 50
