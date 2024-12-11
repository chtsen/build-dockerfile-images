FROM alpine:3.18.3   AS go_build

WORKDIR /opt/yozo/workdir

COPY ./tools/golang/gotpl.go .

RUN set -ex && \
    apk update && \
    apk add --no-cache tzdata  bash bash-doc bash-completion go upx && \
    go mod init gotpl.go  && \
    go build -o gotpl . && \
    upx -9 gotpl 

FROM ubuntu:jammy-20240911.1

LABEL maintainer="Chtsen cyc@yozosoft.com"

COPY ./tools/tpl/yozo-tomcat/server.xml.tpl /usr/local/tomcat/conf/
COPY --from=go_build /opt/yozo/workdir/gotpl /opt/yozo/bin/

ARG TARGETARCH
ARG TOMCAT_V

ENV OS_ARCH="${TARGETARCH:-amd64}" \
    TOMCAT_V="${TOMCAT_V:-9.0.97}" \
    TZ=Asia/Shanghai \
    APP_PORT=8080 \
    LANG=zh_CN.utf8 \
    LC_ALL=zh_CN.utf8 \
    SHUTDOWN_PORT=8005 

# 设置用户/组
RUN set -ex && \
    groupadd -g 1099 yozo && \
    useradd -m -g 1099 -u 1099 yozo && \
    apt update && \
    apt install sudo language-pack-zh-han* libxrender1 libxext6 libssl-dev curl -y && \
    echo "yozo ALL=(ALL)       NOPASSWD:ALL" >> /etc/sudoers && \
    curl -u temp:releases@2022 -SsLf -O "https://releases.yozosoft.com:4431/fcs/tools/jdk/YOZO_Jre_${OS_ARCH}_1.8.0.tar.gz" -O "https://releases.yozosoft.com:4431/fcs/tools/jdk/OpenJDK8U-jdk_${OS_ARCH}_linux_hotspot_8u402b06.tar.gz" -O "https://releases.yozosoft.com:4431/fcs/tools/tomcat/apache-tomcat-${TOMCAT_V}.tar.gz"  && \
    mkdir -p /opt/yozo/yozo_office_jre/ /opt/yozo/open_jdk /opt/yozo/fcscloud  && \
    tar -xf "YOZO_Jre_${OS_ARCH}_1.8.0.tar.gz" -C /opt/yozo/yozo_office_jre/ --no-same-owner && \
    tar -xf "OpenJDK8U-jdk_${OS_ARCH}_linux_hotspot_8u402b06.tar.gz" -C /opt/yozo/open_jdk --strip-components=1 && \
    tar -xf apache-tomcat-9.0.86.tar.gz -C /opt/yozo/fcscloud --strip-components=1 && \
    sed -i '125i\export JAVA_HOME=/opt/yozo/open_jdk /opt/yozo/fcscloud/bin/catalina.sh && \
    rm -rf "YOZO_Jre_${OS_ARCH}_1.8.0.tar.gz"
    #echo "yozo ALL=(ALL)       NOPASSWD:/usr/bin/cat /sys/class/dmi/id/product_uuid,/bin/cat /sys/class/dmi/id/product_uuid" >> /etc/sudoers

ENV TZ=Asia/Shanghai
