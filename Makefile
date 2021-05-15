# get variable from env file
include app.env
# environment variables
NAME					= $(APP_NAME)
VERSION 				= $(shell git describe --tags --always)

BUILD_LDFLAGS 			= -s -w
DOCKER_CONTAINER_NAME	= $(NAME)
DOCKER_IMAGE_NAME		= $(NAME):$(VERSION)

IMAGE_NAME_EXIST 		= $(shell docker images -aq ${DOCKER_IMAGE_NAME})
CONTAINER_NAME_EXIST 	= $(shell docker ps -aq --filter name=${DOCKER_CONTAINER_NAME})
SYSTEM_NAME				= ${NAME}" - "${VERSION}


# ┌┬┐┌─┐┌─┐┌┬┐
#  │ ├┤ └─┐ │ 
#  ┴ └─┘└─┘ ┴ 

tidy:
	@echo "Running $@"
	@go mod tidy

lint:
	@clear
	@go fmt ./...

# ┌┐ ┬ ┬┬┬  ┌┬┐
# ├┴┐│ │││   ││
# └─┘└─┘┴┴─┘─┴┘
build:
	@echo "Building flottbot binary to './binary'"
	@go build -a \
		-ldflags '$(BUILD_LDFLAGS)' -o $(PWD)/binary ./cmd


# ┌┬┐┌─┐┌─┐┬┌─┌─┐┬─┐
#  │││ ││  ├┴┐├┤ ├┬┘
# ─┴┘└─┘└─┘┴ ┴└─┘┴└─

#  ___   ___   ___ _  _____ ___ 
#  |   \ / _ \ / __| |/ / __| _ \
#  | |) | (_) | (__| ' <| _||   /
#  |___/ \___/ \___|_|\_\___|_|_\
                               
docker-build:
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .
	# @docker image prune --filter label=tagged=builder-${DOCKER_IMAGE_NAME} --force


docker-build_second:
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile .

docker-base:
	@echo "Creating base $@ image"
	@docker build \
		-f "./Dockerfile" \
		-t $(DOCKER_IMAGE_NAMEX):$(VERSION) \
		-t $(DOCKER_IMAGE_NAMEX):latest .


docker-flavors:
	@echo "Creating image for $$flavor"; 
	@docker build \
		--build-arg "VERSION=$(VERSION)" \
		--build-arg "GIT_HASH=$(GIT_HASH)" \
		-f "./docker/Dockerfile.$$flavor" \
		-t $(DOCKER_IMAGE_NAME):$$flavor \
		-t $(DOCKER_IMAGE_NAME):$$flavor-$(VERSION) .; \

docker-clear:
	@echo "Cleaning Docker kawalrencanamu";
	@if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;

compose-build:
	@echo "Compose Up";
	@docker-compose up --build;

compose-clean:
	@clear
	# @echo "Compose Clean - "$(SYSTEM_NAME);
	@echo "Compose Clean";
	@docker-compose stop
	@if [ -n "$(CONTAINER_NAME_EXIST)" ]; then docker rm $(CONTAINER_NAME_EXIST) --force; fi;
	@if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;



