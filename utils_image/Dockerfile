FROM ubuntu:bionic

RUN apt-get update \
  && apt-get install --no-install-recommends --yes --force-yes \
    locales \
    curl \
    dnsutils \
    httpie \
    iputils-ping \
    net-tools \
    telnet \
    vim \
    wget \
    netcat-openbsd \
    iproute2

RUN locale-gen en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

EXPOSE 8888
