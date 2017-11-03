FROM golang

# RUN set -ex \
#   && apt-get update && apt-get install -y \
#     bash \
#   && rm -rf /var/lib/apt/lists/*

RUN set -ex \
  && go get -u github.com/labstack/echo/...

RUN set -ex \
  && go get github.com/lib/pq \
  && go get github.com/jmoiron/sqlx \
  && go get github.com/codegangsta/gin

ENV BIN_APP_PORT 1323
