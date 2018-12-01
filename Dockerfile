FROM golang:1.10.3-alpine

ENV GOBIN /go/bin

ADD . /go/src/app
WORKDIR /go/src/app

RUN apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  # ホットリロード
  && go get -u github.com/codegangsta/gin \
  # goデバッグツール
  && go get -u github.com/derekparker/delve/cmd/dlv \
  && dep ensure
