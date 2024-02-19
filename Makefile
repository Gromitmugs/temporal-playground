.PHONY: g generate
g generate:
	make -C thirdparty g

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
up:
	$(MAKE) build
	$(MAKE) build-thirdparty
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

build: $(APP_BIN)
$(APP_BIN): $(shell find build/dockerfile -type f) $(shell find job -type f) $(shell find service -type f) $(shell find thirdparty/client -type f) go.mod go.sum main.go
	go build -o $(APP_BIN)
	docker build -t temporal-worker -f ./build/dockerfile/Dockerfile-worker .
	docker build -t temporal-builder-worker -f ./build/dockerfile/Dockerfile-builder --target kaniko .

build-thirdparty:
	make -C thirdparty build
