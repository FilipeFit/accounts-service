FROM golang:1.16-alpine3.13 AS builder
MAINTAINER Filipe Torqueto

ENV PROFILE=release
ENV APP_PORT=5006

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

EXPOSE $PORT
CMD ["/build/account-service"]