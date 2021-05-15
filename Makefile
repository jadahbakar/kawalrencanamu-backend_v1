#  ____ __ _  _  _ __ ____  __  __ _ _  _ ____ __ _ ____ 
# (  __|  ( \/ )( (  |  _ \/  \(  ( ( \/ |  __|  ( (_  _)
#  ) _)/    /\ \/ /)( )   (  O )    / \/ \) _)/    / )(  
# (____)_)__) \__/(__|__\_)\__/\_)__)_)(_(____)_)__)(__) 
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


#  ____ ____ ____ ____ 
# (_  _|  __) ___|_  _)
#   )(  ) _)\___ \ )(  
#  (__)(____|____/(__) 

tidy:
	@echo "Running $@"
	@go mod tidy

lint:
	@clear
	@go fmt ./...

#  ____  _  _ __ __   ____ 
# (  _ \/ )( (  |  ) (    \
#  ) _ () \/ ()(/ (_/\) D (
# (____/\____(__)____(____/

build:
	@echo "Building flottbot binary to './binary'"
	@go build -a \
		-ldflags '$(BUILD_LDFLAGS)' -o $(PWD)/binary ./cmd



#  ____  __   ___ __ _ ____ ____ 
# (    \/  \ / __|  / |  __|  _ \
#  ) D (  O | (__ )  ( ) _) )   /
# (____/\__/ \___|__\_|____|__\_)
                                          
                               
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

#  ____  __   ___ __ _ ____ ____      ___ __  _  _ ____  __  ____ ____ 
# (    \/  \ / __|  / |  __|  _ \___ / __)  \( \/ |  _ \/  \/ ___|  __)
#  ) D (  O | (__ )  ( ) _) )   (___| (_(  O ) \/ \) __(  O )___ \) _) 
# (____/\__/ \___|__\_|____|__\_)    \___)__/\_)(_(__)  \__/(____(____)

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



