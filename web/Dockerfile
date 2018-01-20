FROM nginx

COPY ./web-main/dist /usr/share/nginx/web-main
COPY ./web-operation/dist /usr/share/nginx/web-operation

COPY ./web/nginx.conf /etc/nginx/nginx.conf
COPY ./web/nginx-web-main.conf /etc/nginx/conf.d/default.conf
COPY ./web/nginx-web-operation.conf /etc/nginx/conf.d/web-operation.conf
COPY ./web/nginx-web-api.conf /etc/nginx/conf.d/api.conf
