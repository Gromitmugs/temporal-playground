.PHONY: update-submodules
update-submodules:
	git submodule update --recursive --init


### Docker Compose ###
export COMPOSE_PROJECT_NAME:=temporal
export ELASTICSEARCH_VERSION:=7.16.2
export POSTGRESQL_VERSION:=13
export TEMPORAL_VERSION:=1.22.0
export TEMPORAL_UI_VERSION:=2.10.3

.PHONY: up
up:
	docker compose up
