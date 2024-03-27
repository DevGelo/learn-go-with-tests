// https://go.dev/play/

package main

import (
	"testing"
)

// Sum calculates the total from an array of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func TestSum(t *testing.T) {
	t.Run("Sum collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
