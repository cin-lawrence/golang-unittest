package inmemory_test

import (
	"testing"

	"ch24/fakes/adapters/inmemory"
	"ch24/fakes/domain/planner"
)

func TestInMemoryAPI1(t *testing.T) {
	(&planner.API1Contract{
		NewAPI1: func() planner.API1 {
			return inmemory.NewAPI1()
		},
	}).Test(t)
}
