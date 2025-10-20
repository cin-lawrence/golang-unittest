package interactions_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"ch23/greet/domain/interactions"
	"ch23/greet/specs"
)

func TestGreet(t *testing.T) {
	specs.AssertGreetSpecification(
		t,
		specs.GreetAdapter(interactions.Greet),
	)

	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
