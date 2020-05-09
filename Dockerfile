FROM golang:alpine3.11 as BuildImage

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add --no-cache gcc musl-dev

ENV GOPATH=/gopath \
    GOBIN=/gopath/bin \
    GOCACHE=/tmp/.gocache \
    PROJPATH=/gopath/src/github.com/ns-stored

COPY . /gopath/src/github.com/ns-stored
WORKDIR /gopath/src/github.com/ns-stored

RUN go build -mod vendor -o bin/captcha-service .

FROM alpine:3.11