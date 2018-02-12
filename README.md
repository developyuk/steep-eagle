
```
         __                                      __   
   _____/ /____  ___  ____     ___  ____ _____ _/ /__ 
  / ___/ __/ _ \/ _ \/ __ \   / _ \/ __ `/ __ `/ / _ \
 (__  ) /_/  __/  __/ /_/ /  /  __/ /_/ / /_/ / /  __/
/____/\__/\___/\___/ .___/   \___/\__,_/\__, /_/\___/ 
                  /_/                  /____/
```
description
# How to run

clone this project with

```
$ git clone https://github.com/developyuk/steep-eagle.git
```

or download this project

```
$ wget https://github.com/developyuk/steep-eagle/archive/develop.zip
```

make sure you have install [docker compose](https://docs.docker.com/compose/install/#install-compose). then

```
$ cd /path/to/project
$ docker-compose build
$ docker-compose up
```

# Develop

## Update schema.sql

After updating schema.sql, you need to [remove postgres container]((https://gist.github.com/bastman/5b57ddb3c11942094f8d0a97d461b430)) to reset database schema

make sure you have exits postgres container first

```
$ cd /path/to/project
$ docker-compose down
$ docker container prune
```

# Powered by

- [Echo](https://echo.labstack.com/) High performance, extensible, minimalist Go web framework
- [PostgreSQL](https://www.postgresql.org/) The world's most advanced open source database
- [Vue.js](https://vuejs.org/) The Progressive JavaScript Framework
- [Material Components for the Web](https://material.io/components/web/) Create beautiful apps with modular and customizable UI components.
- [Normalize.css](https://necolas.github.io/normalize.css/) A modern, HTML5-ready alternative to CSS resets
- [Moment.js](http://momentjs.com/) Parse, validate, manipulate, and display dates and times in JavaScript.
- [Axios](https://github.com/axios/axios) Promise based HTTP client for the browser and node.js
- [Sass](http://sass-lang.com/) CSS with superpowers
- [Pug](https://pugjs.org/api/getting-started.html) Robust, elegant, feature rich template engine for Node.js 
- [Webpack](https://webpack.js.org/) Open-source JavaScript module bundler
- [Imagesweserv](https://images.weserv.nl/) Image cache & resize proxy
- [Docker](https://www.docker.com/) Operating-system-level virtualization also known as containers
