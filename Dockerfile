FROM golang:1.18

WORKDIR /go/src

COPY . /go/src

RUN apt update && apt install -y redis-server

RUN make install

ENTRYPOINT redis-server --protected-mode no