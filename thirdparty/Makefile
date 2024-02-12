.PHONY: g generate
g generate: gen-app gen-client

.PHONY: gen-app
gen-app:
	gqlgen generate --config codegen/gqlgen.yml

.PHONY: c clean
c clean:
	rm -rf internal/graph/gqlmodel/models_gen.go
	rm -rf internal/graph/gqlgen

gen-client-schema:
	for f in graphql/*; do cat $$f; echo; done > client/codegen/schema.graphql

gen-client: gen-client-schema
	cd client/codegen && genqlient

APP_BIN := ./bin/message
build: generate
	go build -o ./bin/message

docker-build: build
	docker build -t message .