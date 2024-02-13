.PHONY: update-submodules
update-submodules:
	git submodule update --recursive --init


### Docker Compose ###
export COMPOSE_PROJECT_NAME=temporal
export CASSANDRA_VERSION=3.11.9
export ELASTICSEARCH_VERSION=7.16.2
export MYSQL_VERSION=8
export TEMPORAL_VERSION=1.22.4
export TEMPORAL_UI_VERSION=2.22.3
export POSTGRESQL_VERSION=13
export POSTGRES_PASSWORD=temporal
export POSTGRES_USER=temporal
export POSTGRES_DEFAULT_PORT=5432
export OPENSEARCH_VERSION=2.5.0

.PHONY: up
up: docker-build-worker docker-build-thirdparty
	docker compose up

.PHONY: helloworld-worker
helloworld-worker:
	go run job/helloworld/worker/main.go

.PHONY: helloworld-starter
helloworld-starter:
	go run job/helloworld/starter/main.go

.PHONY: broadcast-starter
broadcast-starter:
	go run main.go broadcast


APP_BIN := ./build/bin/worker
.PHONY: build
build:
	go build -o $(APP_BIN)

docker-build-worker: build
	docker build -t temporal-worker -f ./build/dockerfile/Dockerfile-worker .

docker-build-thirdparty:
	make -C thirdparty docker-build
