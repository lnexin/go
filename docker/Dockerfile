FROM ubuntu:22.04
MAINTAINER xind "xindu.zh@gmail.com"
RUN apt-get -y update && apt-get install -y nginx
COPY html /usr/share/nginx/html
EXPOSE 80
STOPSIGNAL SIGQUIT

