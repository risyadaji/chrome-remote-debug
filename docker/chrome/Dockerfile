FROM ubuntu:16.04

ARG DEBIAN_FRONTEND=noninteractive

RUN	apt-get update \
	&& apt-get install -y --no-install-recommends gettext-base xvfb fonts-takao supervisor x11vnc socat libxml2 cpp x11-utils x11-xserver-utils xml-core dbus-x11 \
	&& apt-get clean \
	&& rm -rf /var/cache/* /var/log/apt/* /tmp/*

COPY etc/apt /etc/apt
COPY tmp /tmp

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

COPY etc/opt/chrome/policies/managed /etc/opt/chrome/policies/managed
COPY supervisord.conf.template /
COPY run.sh /
CMD ["/run.sh"]
