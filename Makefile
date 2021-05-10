REGISTRY?=registry.grofer.io/ci
NAME?=flagger-slack-handler
TAG?=master-$(shell git rev-parse --short=7 HEAD)
IMAGE_NAME?=$(REGISTRY)/$(NAME):$(TAG)

build:
	CGO_ENABLED=0 go build -a -o ./bin/$(NAME) ./cmd

docker_build:
	docker build -t $(IMAGE_NAME) .

docker_push:
	docker push $(IMAGE_NAME)

docker: docker_build docker_push

