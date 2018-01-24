FROM ubuntu:16.04

ARG DEBIAN_FRONTEND=noninteractive

RUN	apt-get update \
	&& apt-get install -y --no-install-recommends gettext-base xvfb fonts-takao supervisor x11vnc socat libxml2 cpp x11-utils x11-xserver-utils xml-core dbus-x11 \
	&& apt-get clean \
	&& rm -rf /var/cache/* /var/log/apt/* /tmp/*

COPY etc/apt /etc/apt
COPY tmp /tmp

ENV LAST_CHROME_UPDATE 2018-01-24
# Copied the signing key from https://dl.google.com/linux/linux_signing_key.pub
RUN	apt-key add /tmp/linux_signing_key.pub \
	&& apt-get update \
	&& apt-get install -y google-chrome-stable \
	&& apt-get clean \
	&& rm -rf /var/cache/* /var/log/apt/* /tmp/*

RUN addgroup chrome-remote-debugging \
	&& useradd -m -G chrome-remote-debugging chrome

VOLUME ["/home/chrome"]

EXPOSE 5900
EXPOSE 9222

COPY supervisord.conf.template /
COPY run.sh /
CMD ["/run.sh"]
