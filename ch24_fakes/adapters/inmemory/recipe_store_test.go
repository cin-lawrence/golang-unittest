package inmemory_test

import (
	"testing"

	"ch24/fakes/adapters/inmemory"
	"ch24/fakes/domain/planner"
)

func TestInMemoryRecipeStore(t *testing.T) {
	(&planner.RecipeBookContract{NewBook: func() planner.RecipeBook {
		return inmemory.NewRecipeStore()
	}}).Test(t)
}
