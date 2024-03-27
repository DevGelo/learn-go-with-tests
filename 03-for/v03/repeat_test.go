// https://go.dev/play/

package main

import (
	"testing"
)

// Repeat return character repeated repeatCount times.
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}

func TestRepeat(t *testing.T) {
	t.Run("Repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Repeat 'b' 1 times", func(t *testing.T) {
		repeated := Repeat("b", 1)
		expected := "b"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Repeat 'd' -5 times", func(t *testing.T) {
		repeated := Repeat("d", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Repeat 'c' 0 times", func(t *testing.T) {
		repeated := Repeat("c", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
