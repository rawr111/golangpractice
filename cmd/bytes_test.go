package main

import "testing"

var result string

func BenchmarkConversion(b *testing.B) {

	data := []byte("hello world")
	for b.Loop() {
		result = string(data)
	}

}

func BenchmarkNoConversion(b *testing.B) {
	data := []byte("hello world")

	for b.Loop() {
		_ = data
	}
}

func BenchmarkMap1(b *testing.B) {
	mymap := map[int]int{}

	var i int
	for b.Loop() {
		mymap[i] = i
		i++
	}
}

func BenchmarkMap2(b *testing.B) {
	mymap := make(map[int]int, 20880519)

	var i int
	for b.Loop() {
		mymap[i] = i
		i++
	}
}
