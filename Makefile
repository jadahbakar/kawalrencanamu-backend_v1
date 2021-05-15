# environment variables

BUILD_LDFLAGS 			= -s -w
VERSION 				= $(shell git describe --tags --always)
DOCKER_IMAGE 			?= "slackman/kawalrencanamu"
IMAGE_NAME_EXIST 		= $(shell docker images -aq ${DOCKER_IMAGE_NAME})
CONTAINER_NAME_EXIST 	= $(shell docker ps -aq --filter name=${DOCKER_CONTAINER_NAME})
APPLICATION_NAME		= "Kawal Rencanamu"
SYSTEM_NAME				= ${APPLICATION_NAME}"- ver. "${VERSION}

all: test build


# ┌┐ ┬ ┬┬┬  ┌┬┐
# ├┴┐│ │││   ││
# └─┘└─┘┴┴─┘─┴┘
.PHONY: build
build: clean
	@echo "Building flottbot binary to './binary'"
	@go build -a \
		-ldflags '$(BUILD_LDFLAGS)' -o $(PWD)/binary ./cmd


# ┌┬┐┌─┐┌─┐┬┌─┌─┐┬─┐
#  │││ ││  ├┴┐├┤ ├┬┘
# ─┴┘└─┘└─┘┴ ┴└─┘┴└─
docker-base:
	@echo "Creating base $@ image"
	@docker build \
		--build-arg "VERSION=$(VERSION)" \
		--build-arg "GIT_HASH=$(GIT_HASH)" \
		-f "./Dockerfile" \
		-t $(DOCKER_IMAGE):$(VERSION) \
		-t $(DOCKER_IMAGE):latest .


docker-flavors:
	@echo "Creating image for $$flavor"; 
	@docker build \
		--build-arg "VERSION=$(VERSION)" \
		--build-arg "GIT_HASH=$(GIT_HASH)" \
		-f "./docker/Dockerfile.$$flavor" \
		-t $(DOCKER_IMAGE):$$flavor \
		-t $(DOCKER_IMAGE):$$flavor-$(VERSION) .; \

docker-clear:
	@echo "Cleaning Docker kawalrencanamu";
	@if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;

compose-build:
	@echo "Compose Up";
	@docker-compose up --build;

compose-clean:
	@clear
	@echo "Compose Clean - "$(SYSTEM_NAME) ;
	@docker-compose stop
	@if [ -n "$(CONTAINER_NAME_EXIST)" ]; then docker rm $(CONTAINER_NAME_EXIST) --force; fi;
	@if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;



