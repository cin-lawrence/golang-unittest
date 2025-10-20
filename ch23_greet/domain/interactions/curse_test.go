package interactions_test

import (
	"testing"

	"ch23/greet/domain/interactions"
	"ch23/greet/specs"
)

func TestCurseAdapter(t *testing.T) {
	specs.AssertCurseSpecification(
		t,
		specs.CurseAdapter(interactions.Curse),
	)
}
