FROM golang:1.7.3

MAINTAINER Zonesan <chaizs@asiainfo.com>

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY . /go/src/github.com/asiainfoLDP/datafoundry-gitter

WORKDIR /go/src/github.com/asiainfoLDP/datafoundry-gitter

EXPOSE 7000

RUN go build

ENTRYPOINT ["./datafoundry-gitter"]







