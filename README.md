
# Steep Eagle
[![Build Status](https://travis-ci.org/developyuk/steep-eagle.svg?branch=master)](https://travis-ci.org/developyuk/steep-eagle)

_description_

## Instalation

Clone this project with
```
$ git clone https://github.com/developyuk/steep-eagle.git
```
or download

```
$ wget https://github.com/developyuk/steep-eagle/archive/develop.zip
```

Make sure you have install [docker compose](https://docs.docker.com/compose/install/#install-compose).
```
docker -v
docker-compose -v
```
then run with
```
$ cd /path/to/project
$ ./scripts/up-dev.sh
```

## Develop
### Updating database schema

After updating schema.sql, you need to [remove postgres container]((https://gist.github.com/bastman/5b57ddb3c11942094f8d0a97d461b430)) to reset database schema

make sure you have stop postgres container first
```
$ cd /path/to/project
$ docker-compose down
$ docker container prune
```

## Powered by
- [Docker](//www.docker.com/) Operating-system-level virtualization also known as containers
- [Echo](//echo.labstack.com/) High performance, extensible, minimalist Go web framework
- [PostgreSQL](//www.postgresql.org/) The world's most advanced open source database
- [PostgREST](//postgrest.com/) PostgREST is a standalone web server that turns your PostgreSQL database directly into a RESTful API.
- [Vue.js](//vuejs.org/) The Progressive JavaScript Framework
- [Webpack](//webpack.js.org/) Open-source JavaScript module bundler
- [Axios](//github.com/axios/axios) Promise based HTTP client for the browser and node.js
- [Sass](//sass-lang.com/) CSS with superpowers
- [Pug](//pugjs.org/) Robust, elegant, feature rich template engine for Node.js 
- [Lodash](//lodash.com/) A modern JavaScript utility library delivering modularity, performance & extras.
- [Normalize.css](//necolas.github.io/normalize.css/) A modern, HTML5-ready alternative to CSS resets
- [Material Components for the Web](//material.io/components/web/) Create beautiful apps with modular and customizable UI components.
- [hammerjs](//hammerjs.github.io/) Add touch gestures to your webapp.
- [Moment.js](//momentjs.com/) Parse, validate, manipulate, and display dates and times in JavaScript.
- [Imagesweserv](//images.weserv.nl/) Image cache & resize proxy
