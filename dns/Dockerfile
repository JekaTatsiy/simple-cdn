FROM ubuntu:latest

RUN apt-get update \
  && apt-get install -y \
  bind9 \
  bind9utils \
  bind9-doc \
  dnsutils \
  iputils-ping \
  net-tools \
  curl \ 
  systemd \
  vim

RUN service named start

# Enable IPv4
# RUN sed -i 's/OPTIONS=.*/OPTIONS="-4 -u bind"/' /etc/default/bind9

# Copy configuration files
COPY bind /etc/bind

# Run eternal loop
CMD ["named", "-gu","bind"]