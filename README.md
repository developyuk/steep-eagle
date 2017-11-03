
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
$ docker rm $(docker ps -qa --no-trunc --filter "status=exited")
```

# Powered by

- [vue.js](https://vuejs.org/)
- [golang echo](https://echo.labstack.com/)
- [postgresql](https://www.postgresql.org/)
