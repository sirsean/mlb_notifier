FROM debian:latest

MAINTAINER Sean Schulte

RUN apt-get update --yes
RUN apt-get dist-upgrade --yes

RUN apt-get install --yes \
    git \
    golang

ENV GOPATH /

ADD . /src/github.com/sirsean/mlb_notifier
WORKDIR /src/github.com/sirsean/mlb_notifier

RUN go get ./...
RUN go build

ENTRYPOINT ["./mlb_notifier"]
