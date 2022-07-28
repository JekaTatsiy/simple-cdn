FROM nginx:latest

RUN  apt-get update && apt-get install -y vim dnsutils curl net-tools

COPY nginx.conf /etc/nginx
COPY src /etc/nginx/data/src