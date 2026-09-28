package utils

import (
	"math/rand"
	"slices"
	"strings"
)

// удаляет элемент из определенного индекса слайса
func DeleteElementFromSlice[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// вставляет элемент в нужный индекс слайса
func InsertElementToSlice[T comparable](slice []T, index int, element T) []T {
	var emptyEl T
	slice = append(slice, emptyEl)
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
}

// удаляет последний элемент слайса
func Pop[T comparable](slice []T) ([]T, T) {
	return slice[:len(slice)-1], slice[len(slice)-1]
}

// перемешивает элементы слайса
func ShuffleSlice[T comparable](slice []T) []T {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// разбивает слайс на батчи заданного размера
func BatchSlice[T comparable](slice []T, batchSize int) [][]T {
	batches := make([][]T, 0, (len(slice)+batchSize-1)/batchSize)
	for batchSize < len(slice) {
		slice, batches = slice[batchSize:], append(batches, slice[0:batchSize:batchSize])
	}
	batches = append(batches, slice)
	return batches
}

// фильтрует слайс без реаллокации
func FilterSlice[T comparable](slice []T, predicat func(x T) bool) []T {
	n := 0
	for _, a := range slice {
		if predicat(a) {
			slice[n] = a
			n++
		}
	}
	return slice[:n]
}

func RemoveDublicates[T comparable](slice []T) []T {
	elementSet := make(map[T]struct{}, len(slice))
	result := make([]T, 0, len(slice))

	for _, a := range slice {
		if _, exist := elementSet[a]; !exist {
			elementSet[a] = struct{}{}
			result = append(result, a)
		}
	}

	return result
}

func ReverseSlice[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}

func ConcatSlices[T comparable](s ...[]T) []T {
	return slices.Concat(s...)
}

func GroupBy[T comparable, E comparable](slice []T, p func(el T) E) map[T][]E {
	result := make(map[T][]E, len(slice))

	for _, x := range slice {
		result[x] = append(result[x], p(x))
	}

	return result
}

func JoinWords(words []string) string {
	var b = strings.Builder{}

	var totalLength int
	for _, word := range words {
		totalLength += len(word)
	}
	if len(words) > 0 {
		totalLength += len(words) - 1
	}

	b.Grow(totalLength)
	for _, word := range words {
		b.WriteString(word)
		b.WriteByte(' ')
	}
	return b.String()
}
