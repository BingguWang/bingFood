FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

MAINTAINER "bingguWang@441282413@qq.com"

#WORKDIR /go/src/app

COPY --from=0 /go/src/app/server ./
COPY --from=0 /go/src/app/config.yaml ./

EXPOSE 8088
ENTRYPOINT ./server
