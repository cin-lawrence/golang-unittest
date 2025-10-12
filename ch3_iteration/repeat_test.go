package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat with default count", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})

	t.Run("repeat with any count less than 0", func(t *testing.T) {
		repeated := Repeat("a", -999)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})

	t.Run("repeat with count = 10", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}
