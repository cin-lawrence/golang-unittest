package sqlite_test

import (
	"testing"

	"ch24/fakes/adapters/sqlite"
	"ch24/fakes/domain/planner"
)

func TestRecipeStore(t *testing.T) {
	(&planner.RecipeBookContract{
		NewBook: func() planner.RecipeBook {
			return sqlite.NewRecipeStore(sqlite.NewSQLiteClient())
		},
	}).Test(t)
}
