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

func GroupBy[T comparable, GroupByParam comparable, GroupByValue comparable](slice []T, p func(el T) (GroupByParam, GroupByValue)) map[GroupByParam][]GroupByValue {
	result := make(map[GroupByParam][]GroupByValue, len(slice))

	for _, x := range slice {
		param, value := p(x)
		result[param] = append(result[param], value)
	}

	return result
}

func CountSliceElementsFrequency[T comparable](slice []T) map[T]int {
	result := make(map[T]int, len(slice))
	for _, el := range slice {
		if _, ok := result[el]; ok {
			result[el]++
		} else {
			result[el] = 1
		}
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

func Split(input string, key string) []string {
	if key == "" {
		return []string{input}
	}

	result := []string{}
	inputRunes := []rune(input)
	keyRunes := []rune(key)
	word := strings.Builder{}
	word.Grow(len(input))

	var i int = 0
	for i < len(inputRunes) {
		var isCorrectSeparator bool = true
		for j, keyRune := range keyRunes {
			if i+j >= len(inputRunes) || keyRune != inputRunes[i+j] {
				isCorrectSeparator = false
				break
			}
		}

		if isCorrectSeparator {
			result = append(result, word.String())
			word.Reset()
			i += len(keyRunes)
		} else {
			word.WriteRune(inputRunes[i])
			i++
		}
	}

	result = append(result, word.String())

	return result
}

const lettersDiff = 'a' - 'A'

// капитализирует все буквы латинского алфавита в строке
func CapitalizeLetters(s string) string {
	newStr := strings.Builder{}
	newStr.Grow(len(s))

	for _, r := range s {
		var newRune rune = r
		if r >= 'a' && r <= 'z' {
			newRune -= lettersDiff
		}
		newStr.WriteRune(newRune)
	}

	return newStr.String()
}
