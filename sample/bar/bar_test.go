package bar_test

import (
	"testing"

	"lawrence/sample/bar"
)

func TestSomething(t *testing.T) {
	t.Run("test simple add", func(t *testing.T) {
		got := bar.Something(2, 3)
		want := 5

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
