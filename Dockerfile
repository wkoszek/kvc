FROM golang:1.18

WORKDIR /go/src

COPY . /go/src

RUN apt update && apt install -y redis-server && make install

ENTRYPOINT redis-server --protected-mode no