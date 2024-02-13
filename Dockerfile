FROM golang:1.21.4

WORKDIR /go/app

COPY ./build/bin/worker /go/app/worker

EXPOSE 8001
