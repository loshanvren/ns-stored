FROM golang:alpine3.11 as BuildImage

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add --no-cache gcc musl-dev

ENV GOPATH=/gopath \
    GOBIN=/gopath/bin \
    GOCACHE=/tmp/.gocache \
    PROJPATH=/gopath/src/github.com/ns-stored

COPY . /gopath/src/github.com/ns-stored
WORKDIR /gopath/src/github.com/ns-stored

RUN go build -mod vendor -o bin/ns-stored .

FROM alpine:3.11

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add --no-cache tcpdump lsof net-tools tzdata curl

ENV TZ=Asia/Shanghai PATH=$PATH:/opt/ns-stored/bin

WORKDIR /opt/ns-stored/bin

COPY --from=BuildImage /gopath/src/github.com/ns-stored/bin/ns-stored /opt/ns-stored/bin/
COPY --from=BuildImage /gopath/src/github.com/ns-stored/env.yaml /opt/ns-stored/
COPY --from=BuildImage /gopath/src/github.com/ns-stored/asset /opt/ns-stored/asset

RUN chmod +x /opt/ns-stored/bin/ns-stored

CMD /opt/ns-stored/bin/ns-stored
