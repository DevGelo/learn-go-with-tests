// https://go.dev/play/

package main

import (
	"testing"
)

// Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func TestSum(t *testing.T) {
	numbers := [...]int{1, 2, 3, 5, 7}

	got := Sum(numbers)
	want := 18

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
