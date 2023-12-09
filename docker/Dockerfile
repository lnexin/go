FROM ubuntu:latest
MAINTAINER xind "xindu.zh@gmail.com"
# RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y nginx
COPY html /usr/share/nginx/html
EXPOSE 80
#STOPSIGNAL SIGQUIT

