.PHONY: g generate
g generate:
	gqlgen generate --config codegen/gqlgen.yml

.PHONY: c clean
c clean:
	rm -rf internal/graph/gqlmodel/models_gen.go
	rm -rf internal/graph/gqlgen
