# app engineが対象としているバージョンが1.11 or 1.12+
FROM golang:1.12.17-alpine

RUN mkdir /workdir
WORKDIR /workdir
ADD . /workdir

CMD ["/usr/local/go/bin/go", "run", "/workdir/main.go"]
