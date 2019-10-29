# Copyright 2018 Payfazz Agent Authors.

# The default docker image name
APP_NAME := chrome-remote-debug

CHROME_IMAGE := chrome-remote
# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

app-container:
	@echo "building the container..."
	@docker build --label "version=$(VERSION)" -t ${CHROME_IMAGE}:$(VERSION) -f ./docker/Dockerfile .

push-chrome-fazz:
	@docker tag $(CHROME_IMAGE):$(VERSION) docker.fazzfinancial.com/payfazz/$(APP_NAME)/$(CHROME_IMAGE):$(VERSION)
	@docker push docker.fazzfinancial.com/payfazz/$(APP_NAME)/$(CHROME_IMAGE):$(VERSION)

# before the push, you need to create the ssh tunnel to the registry first:
# $ ssh -L 5000:locahost:5000 root@10.0.125.236 - DEV
# $ ssh -L 5000:locahost:5000 root@10.0.106.58 - PRD
push:
	@docker tag $(IMAGE):$(VERSION) localhost:5000/$(IMAGE):$(VERSION)
	@docker push localhost:5000/$(IMAGE):$(VERSION) 

push-mac:
	@docker tag $(IMAGE):$(VERSION) docker.fazzfinancial.com/payfazz/$(APP_NAME)/$(IMAGE):$(VERSION)
	@docker push docker.fazzfinancial.com/payfazz/$(APP_NAME)/$(IMAGE):$(VERSION)

clean:
	rm -rf .container-* .dockerfile-* .push-*
	rm -rf .go bin
