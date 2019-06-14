FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

WORKDIR /app

ENV GOBIN /go/bin

RUN go get -u github.com/PuerkitoBio/goquery \
    && go get -u github.com/microcosm-cc/bluemonday \
    && go get -u github.com/djimenez/iconv-go \
    && go get -u github.com/jinzhu/gorm \
    && go get -u github.com/go-sql-driver/mysql
