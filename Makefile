# Copyright 2018 Payfazz Agent Authors.

# The default docker image name
APP_NAME := chrome-headfull

CHROME_IMAGE := chrome-remote
# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

app-container:
	@echo "building the container..."
	@docker build --label "version=$(VERSION)" -t ${CHROME_IMAGE}:$(VERSION) -f ./docker/Dockerfile .

push-chrome-fazz:
	@docker tag $(CHROME_IMAGE):$(VERSION) docker.fazzfinancial.com/payfazz/$(APP_NAME)/$(CHROME_IMAGE):$(VERSION)
	@docker push docker.fazzfinancial.com/payfazz/$(APP_NAME)/$(CHROME_IMAGE):$(VERSION)

clean:
	rm -rf .container-* .dockerfile-* .push-*
	rm -rf .go bin
