FROM nginx:latest

RUN  apt-get update && apt-get install -y vim dnsutils curl net-tools

COPY nginx.conf /etc/nginx/nginx.conf
COPY geo/geo.conf /etc/nginx/geo.conf
#COPY site.com.crt /etc/nginx/ssl/site.com.crt
COPY 127.0.0.1.crt /etc/nginx/ssl/127.0.0.1.crt
COPY device.key /etc/nginx/ssl/device.key