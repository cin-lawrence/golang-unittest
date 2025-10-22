package sqlaite_test

import (
	"testing"

	"ch24/fakes/adapters/sqlaite"
	"ch24/fakes/domain/planner"
)

func TestRecipeStore(t *testing.T) {
	(&planner.RecipeBookContract{
		NewBook: func() planner.RecipeBook {
			return sqlaite.NewRecipeStore(sqlaite.NewSQLiteClient())
		},
	}).Test(t)
}
