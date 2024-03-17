package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		assertCorrectValue(t, got, expected, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		assertCorrectValue(t, got, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !slices.Equal(got, expected) {
			t.Errorf("expected %v | got %v", expected, got)
		}
	})

}

func assertCorrectValue(t *testing.T, got int, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %d | got %d | given %v", expected, got, numbers)
	}
}
