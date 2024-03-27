// https://go.dev/play/

package main

import (
	"testing"
)

// Add takes two integers and returns the sum of them
func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
