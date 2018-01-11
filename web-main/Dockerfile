FROM node

RUN set -ex \
  && npm install -g vue-cli

WORKDIR /project
COPY package.json /project/package.json

RUN set -ex \
  && npm install

RUN set -ex \
  && npm install \
    node-sass sass-loader \
    normalize.css \
    jquery \
    materialize-css \
    sass-bem \
    --save-dev

RUN set -ex \
  && npm install \
    axios \
    --save
