FROM daocloud.io/library/centos:centos7.4.1708
MAINTAINER xind "xindu.zh@gmail.com"
# RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
#RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN cp /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.bak
RUN sed -i 's/enabled=1/enabled=0/g' /etc/yum/pluginconf.d/fastestmirror.conf && \
    sed -i 's/enabled=1/enabled=0/g' /etc/yum/pluginconf.d/ovl.conf && \
    sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-Base.repo && \
    sed -i 's/#baseurl/baseurl/g' /etc/yum.repos.d/CentOS-Base.repo && \
    sed -i 's/http:\/\/mirror.centos.org\/centos/http:\/\/mirror.tuna.tsinghua.edu.cn\/centos/g' /etc/yum.repos.d/CentOS-Base.repo && \
    sed -i 's/http:\/\/mirrorlist.centos.org/#http:\/\/mirrorlist.centos.org/g' /etc/yum.repos.d/CentOS-Base.repo

RUN yum install -y epel-release && yum -y update
RUN yum install -y nginx
#RUN mv /etc/yum.repos.d/CentOS-Base.repo.bak /etc/yum.repos.d/CentOS-Base.repo

COPY html /usr/share/nginx/html
EXPOSE 80

STOPSIGNAL SIGQUIT

