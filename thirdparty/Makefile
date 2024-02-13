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

.PHONY: build
build: $(APP_BIN)

$(APP_BIN): $(shell find internal -type f)
	go build -o $(APP_BIN)
	docker build -t message .