package utils

import (
	"testing"
)

func TestFilterSlice(t *testing.T) {
	t.Run("chec something", func(t *testing.T) {
		filtered := FilterSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(x int) bool {
			return x%3 == 0
		})
		expected := []int{3, 6, 9}
		if !equalSlices(filtered, expected) {
			t.Fatalf("slices are different: %d, %d", filtered, expected)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("chec something", func(t *testing.T) {
		reversed := ReverseSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		if !equalSlices(reversed, expected) {
			t.Fatalf("slices are different: reversed: %d, expected: %d", reversed, expected)
		}
	})
}

func equalSlices[T comparable](firstSlice []T, secondSlice []T) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}

	for i := range firstSlice {
		if firstSlice[i] != secondSlice[i] {
			return false
		}
	}
	return true
}
