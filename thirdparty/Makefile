.PHONY: g generate
g generate:
	gqlgen generate --config codegen/gqlgen.yml

.PHONY: c clean
c clean:
	rm -rf internal/graph/gqlmodel/models_gen.go
	rm -rf internal/graph/gqlgen

gen-client-schema:
	for f in graphql/*; do cat $$f; echo; done > client/codegen/schema.graphql

gen-client: gen-client-schema
	cd client/codegen && genqlient