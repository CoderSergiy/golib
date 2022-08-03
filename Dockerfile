FROM golang:1.18
# https://hub.docker.com/_/golang

ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /go/golib
