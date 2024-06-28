FROM openjdk:8u342-jdk-buster

LABEL maintainer="chtsen cyc@live.cn"

RUN set -ex && \
    apt-get update && \
    apt-get install sudo -y && \
    echo "yozo ALL=(ALL)       NOPASSWD:ALL" >> /etc/sudoers && \
    groupadd yozo -g 1099 && \
    useradd yozo -u 1099 -g 1099 -s /bin/bash -m

