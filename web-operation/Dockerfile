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
    pug pug-loader \
    normalize.css \
    sass-bem \
    jquery \
    materialize-css \
    material-components-web \
    fullcalendar \
    jquery \
    moment \
    --save-dev
