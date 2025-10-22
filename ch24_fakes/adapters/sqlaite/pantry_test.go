package sqlaite_test

import (
	"testing"

	"ch24/fakes/adapters/sqlaite"
	"ch24/fakes/domain/planner"
)

func TestSQLitePantry(t *testing.T) {
	client := sqlaite.NewSQLiteClient()
	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	})

	(&planner.PantryContract{
		NewPantry: func() planner.Pantry {
			return sqlaite.NewPantry(client)
		},
	}).Test(t)
}
