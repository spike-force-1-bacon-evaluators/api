.DEFAULT_GOAL := help

PROJECT_NAME := BACON API

.PHONY: help image run prep test

help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

image: ## build docker image and push image to registry
	@docker build -t alesr/bacon-api -f resources/prod/Dockerfile .
	@docker tag alesr/bacon-api alesr/bacon-api:latest
	@docker push alesr/bacon-api

run: ## deploy docker container
	@docker run --rm -d --net baconnet -p 8080:8080 --name bacon-api alesr/bacon-api

prep: ## download dependencies and format code
	@go get -d ./...
	@go fmt ./...

build/osx: prep ## build osx binary
	@go build

build/linux: prep ## build linux binary
	@GOOS=linux go build

test: ## run unit tests
	@docker build -t bacon-api-test -f resources/prod/Dockerfile .
	@docker run --rm bacon-api-test
