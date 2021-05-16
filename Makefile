#  ____ __ _  _  _ __ ____  __  __ _ _  _ ____ __ _ ____ 
# (  __|  ( \/ )( (  |  _ \/  \(  ( ( \/ |  __|  ( (_  _)
#  ) _)/    /\ \/ /)( )   (  O )    / \/ \) _)/    / )(  
# (____)_)__) \__/(__|__\_)\__/\_)__)_)(_(____)_)__)(__) 
# get variable from env file
include app.env
# environment variables
NAME					= $(APP_NAME)
VERSION 				= $(shell git describe --tags --always)

BUILD_LDFLAGS			= -s -w
DOCKER_CONTAINER_NAME	= $(NAME)
DOCKER_IMAGE_NAME		= $(NAME):$(VERSION)

IMAGE_NAME_EXIST		= $(shell docker images -aq ${DOCKER_IMAGE_NAME})
CONTAINER_NAME_EXIST	= $(shell docker ps -aq --filter name=${DOCKER_CONTAINER_NAME})
DOCKER_HUB_REPO			= dedystyawan/kawalrencanamu-backend

#   __   __    __      __ __ _     __  __ _ ____ 
#  / _\ (  )  (  )    (  |  ( \   /  \(  ( (  __)
# /    \/ (_/\/ (_/\   )(/    /  (  O )    /) _) 
# \_/\_/\____/\____/  (__)_)__)   \__/\_)__|____)

# up: lint clean docker-build docker-push docker-run
up: lint clean docker-build compose-up

clean: compose-down docker-remove-container docker-remove-image

#  ____ ____ ____ ____ 
# (_  _|  __) ___|_  _)
#   )(  ) _)\___ \ )(  
#  (__)(____|____/(__) 

tidy:
	@echo "-> Running $@"
	@go mod tidy

lint:
	# @clear
	@echo "-> Running $@"
	@go fmt ./...

start: lint
	# @clear
	@echo "-> Running $@"
	@go run cmd/main.go


#  ____  _  _ __ __   ____ 
# (  _ \/ )( (  |  ) (    \
#  ) _ () \/ ()(/ (_/\) D (
# (____/\____(__)____(____/

go-build:
	@echo "-> Building flottbot binary to './binary'"
	@go build -a \
		-ldflags '$(BUILD_LDFLAGS)' -o $(PWD)/binary ./cmd


#  ____  __   ___ __ _ ____ ____ 
# (    \/  \ / __|  / |  __|  _ \
#  ) D (  O | (__ )  ( ) _) )   /
# (____/\__/ \___|__\_|____|__\_)
                                          
                               
docker-build:
	@echo "-> Running $@"
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .
# 	# @docker image prune --filter label=tagged=builder-${DOCKER_IMAGE_NAME} --force

# docker-build_second:
# 	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile .

# docker-flavors:
# 	@echo "Creating image for $$flavor"; 
# 	@docker build \
# 		--build-arg "VERSION=$(VERSION)" \
# 		--build-arg "GIT_HASH=$(GIT_HASH)" \
# 		-f "./docker/Dockerfile.$$flavor" \
# 		-t $(DOCKER_IMAGE_NAME):$$flavor \
# 		-t $(DOCKER_IMAGE_NAME):$$flavor-$(VERSION) .; \

docker-push:
	@echo "-> Running $@"
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@echo $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):$(VERSION)
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):$(VERSION)
	@docker push $(DOCKER_HUB_REPO)

# docker-run:
# 	@docker run --rm --detach --name $(DOCKER_CONTAINER_NAME) -p $(PORT):$(PORT) $(DOCKER_HUB_REPO)

# docker-build:
# 	@echo "Creating base $@ image";
# 	@docker build \
# 		-f "./Dockerfile" \
# 		-t $(DOCKER_IMAGE_NAME) .

docker-remove-image:
	@echo "-> Running $@";
	@if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;


docker-remove-container:
	@echo "-> Running $@"
	@if [ -n "$(CONTAINER_NAME_EXIST)" ]; then docker rm $(CONTAINER_NAME_EXIST) --force; fi;


#  ____  __   ___ __ _ ____ ____      ___ __  _  _ ____  __  ____ ____ 
# (    \/  \ / __|  / |  __|  _ \___ / __)  \( \/ |  _ \/  \/ ___|  __)
#  ) D (  O | (__ )  ( ) _) )   (___| (_(  O ) \/ \) __(  O )___ \) _) 
# (____/\__/ \___|__\_|____|__\_)    \___)__/\_)(_(__)  \__/(____(____)

compose-up: compose-down
	@echo "-> Running $@";
	@docker-compose up --build;

compose-down:
	@echo "-> Running $@";
	@docker-compose stop;

compose-clean: compose-down
	@echo "-> Running $@";
	@if [ -n "$(CONTAINER_NAME_EXIST)" ]; then docker rm $(CONTAINER_NAME_EXIST) --force; fi;
