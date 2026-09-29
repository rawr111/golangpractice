package utils

import (
	"slices"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	t.Run("check slice filter", func(t *testing.T) {
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
	t.Run("check slice reverse", func(t *testing.T) {
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

type person struct {
	Name string
	City string
}

func TestGroupBy(t *testing.T) {
	t.Run("check group by map", func(t *testing.T) {
		dataset := []person{
			{
				Name: "Анна",
				City: "Москва",
			},
			{
				Name: "Вика",
				City: "Москва",
			},
			{
				Name: "Антон",
				City: "Москва",
			},
			{
				Name: "Абрам",
				City: "Ереван",
			},
			{
				Name: "Арсен",
				City: "Ереван",
			},
			{
				Name: "Николь",
				City: "Нью Йорк",
			},
		}
		groupped := GroupBy(dataset, func(p person) (string, string) {
			return p.City, p.Name
		})

		peopleFromMoscow, ok := groupped["Москва"]
		slices.Sort(peopleFromMoscow)
		expect := []string{"Анна", "Антон", "Вика"}
		slices.Sort(expect)

		if !ok || !equalSlices(peopleFromMoscow, expect) {
			t.Errorf("group by result is wrong. expected: %v, got: %v", expect, peopleFromMoscow)
		}

		peopleFromNewYork, ok := groupped["Нью Йорк"]
		slices.Sort(peopleFromNewYork)
		expect = []string{"Николь"}
		slices.Sort(expect)

		if !ok || !equalSlices(peopleFromNewYork, expect) {
			t.Errorf("group by result is wrong. expected: %v, got: %v", expect, peopleFromNewYork)
		}
	})
}
