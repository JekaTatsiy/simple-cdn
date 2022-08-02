FROM nginx:latest

RUN  apt-get update && apt-get install -y vim dnsutils curl net-tools

COPY nginx.conf /etc/nginx
COPY site.com.crt /etc/nginx/ssl/site.com.crt
COPY device.key /etc/nginx/ssl/device.key