FROM golang:1.10.3-alpine

ENV GOBIN /go/bin

ADD . /go/src/app
WORKDIR /go/src/app

RUN apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

# goデバッグツール
RUN go get -u github.com/derekparker/delve/cmd/dlv

RUN go build -o /go/bin/myapp .

# docker run -v $PWD:/go/src/app -it golang:1.9.2-alpine /bin/sh
# apk add --no-cache git
# go get -u github.com/golang/dep/cmd/dep
# cd /go/src/app
# dep init
# exit
# docker-compose up -d --build
