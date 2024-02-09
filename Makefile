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
up:
	docker compose up
