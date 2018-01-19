FROM postgres

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD mysecretpassword

COPY ./0-main.sql /docker-entrypoint-initdb.d/0-main.sql
COPY ./1-sessions.sql /docker-entrypoint-initdb.d/1-sessions.sql
COPY ./2-privileges.sql /docker-entrypoint-initdb.d/2-privileges.sql
