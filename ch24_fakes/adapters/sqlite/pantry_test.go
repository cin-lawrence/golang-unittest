package sqlite_test

import (
	"testing"

	"ch24/fakes/adapters/sqlite"
	"ch24/fakes/domain/planner"
)

func TestSQLitePantry(t *testing.T) {
	client := sqlite.NewSQLiteClient()
	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	})

	(&planner.PantryContract{
		NewPantry: func() planner.Pantry {
			return sqlite.NewPantry(client)
		},
	}).Test(t)
}
