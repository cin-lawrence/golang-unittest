.PHONY: entgen
entgen:
	@go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./adapters/sqlite/ent/schema

.PHONY: describe
describe:
	@go run -mod=mod entgo.io/ent/cmd/ent describe ./adapters/sqlite/ent/schema
