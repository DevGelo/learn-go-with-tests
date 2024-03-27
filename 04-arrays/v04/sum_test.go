// https://go.dev/play/

package main

import (
	"slices"
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

// SumAll calculates the respective sums of every slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbersToSum := len(numbersToSum)
	sums := make([]int, lengthOfNumbersToSum)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{10, 100, 1000})
	want := []int{3, 1110}

	if !slices.Equal(got, want) { //if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
