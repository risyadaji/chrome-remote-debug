# Copyright 2019 The Payfazz Agent Authors.
#1 install golang
FROM golang:1.13-alpine as builder
WORKDIR /chrome-remote-debug

COPY go.mod . 
COPY go.sum .

RUN apk add --no-cache ca-certificates git

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 go install -installsuffix "static" ./cmd/httpserver/...

#2 Installing chrome
FROM ubuntu:16.04
ARG DEBIAN_FRONTEND=noninteractive

RUN	apt-get update \
	&& apt-get install -y --no-install-recommends gettext-base xvfb fonts-takao supervisor x11vnc socat libxml2 cpp x11-utils x11-xserver-utils xml-core dbus-x11 \
	&& apt-get clean \
	&& rm -rf /var/cache/* /var/log/apt/* /tmp/*

COPY docker/chrome/etc/apt /etc/apt

RUN apt-get -y install wget

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
	&& sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list' \
	&& apt-get update \
	&& apt-get -y install google-chrome-stable \
	&& apt-get clean \
	&& rm -rf /var/cache/* /var/log/apt/* /tmp/*

RUN addgroup chrome-remote-debugging \
	&& useradd -m -G chrome-remote-debugging chrome 

EXPOSE 5900
EXPOSE 9222

RUN apt-get install -y ca-certificates tzdata
COPY --from=builder /go/bin /bin

# copy chrome file
COPY docker/chrome/etc/opt/chrome/policies/managed /etc/opt/chrome/policies/managed
COPY docker/chrome/supervisord.conf.template /
COPY docker/chrome/run.sh /

CMD ["./run.sh", "/bin/httpserver"]
