package sqlite

import (
	"context"
	"log"

	"entgo.io/ent/dialect"

	"ch24/fakes/adapters/sqlite/ent"
)

func NewSQLiteClient() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
