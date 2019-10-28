# Use the following settings to override repo etc
DOCKER_REPO=webiq
DOCKER_IMAGE_NAME=chrome-remote-debug
TAG=latest
ENVFILE=default.env

default: builddocker

clean:
	docker rmi $(DOCKER_IMAGE_NAME):$(TAG)

builddocker:
	docker build --rm=true --tag=$(DOCKER_IMAGE_NAME):$(TAG) .

tag:
	docker tag $(DOCKER_IMAGE_NAME):$(TAG) $(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(TAG)

push: tag
	docker push $(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(TAG)

run: builddocker
	docker run \
		--name chrome \
		--rm \
		--publish 5900:5900 \
		--publish 9222:9222 \
		--shm-size=1g \
		--env-file $(ENVFILE) \
		--volume ${CURDIR}/extensions:/extensions \
		--volume /:/home/chrome \
		$(DOCKER_IMAGE_NAME):$(TAG)

version: builddocker
	docker run \
		--name chrome \
		--rm \
		$(DOCKER_IMAGE_NAME):$(TAG) \
		google-chrome \
		--version
