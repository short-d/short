FROM nginx:1.17.5-alpine

WORKDIR /usr/share/nginx/html

COPY build/ .

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY default.conf /etc/nginx/conf.d/