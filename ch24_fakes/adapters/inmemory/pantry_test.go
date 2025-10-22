package inmemory_test

import (
	"testing"

	"ch24/fakes/adapters/inmemory"
	"ch24/fakes/domain/planner"
)

func TestInMemoryPantry(t *testing.T) {
	(&planner.PantryContract{
		NewPantry: func() planner.Pantry {
			return inmemory.NewPantry()
		},
	}).Test(t)
}
