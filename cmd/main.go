package main

import (
	"fmt"
	"go-base-toolkit/internal/utils"
	"io"
	"sync"
)

type MyReader struct {
	data []byte
	pos  int
}

func (r *MyReader) Read(p []byte) (int, error) {
	if r.pos == len(r.data) {
		return 0, io.EOF
	}

	n := copy(p, r.data[r.pos:])
	r.pos += n

	return n, nil
}

func NewReader(s string) io.Reader {
	return &MyReader{
		data: []byte(s),
		pos:  0,
	}
}

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%c\n", 1071)

	// myReader := NewReader("bebra test")
	// someWriter := strings.Builder{}

	// written, err := io.Copy(&someWriter, myReader)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(written, someWriter.String())
}

func alotofgorutines() {
	m := sync.Map{}
	m.Store("йоу", 1)
	for range 1000 {
		go func() {
			v, _ := m.Load("йоу")
			n, ok := v.(int)
			if !ok {
				panic("")
			}
			n = n + 1
			m.Store("йоу", v)
		}()
	}
	r, _ := m.Load("йоу")
	fmt.Println(r)
}

func alotofgorutines2() {
	m := map[string]int{}
	m["йоу"] = 1
	for range 1000 {
		go func() {
			v, _ := m["йоу"]
			m["йоу"] = v + 1
		}()
	}
	fmt.Println(m["йоу"])
}

func testslice() {
	m := make([]int, 1)
	m[0] = 1
	for range 1000 {
		go func() {
			m[0] += 1
		}()
	}
	fmt.Println(m[0])
}

func testScenario(a []int) {
	a = utils.InsertElementToSlice(a, 4, 9)
	fmt.Println(a)

	a = utils.FilterSlice(a, func(x int) bool {
		return x%2 == 0
	})

	a = utils.ShuffleSlice(a)
	fmt.Println(a)

	a = utils.ReverseSlice(a)
	fmt.Println(a)

	batches := utils.BatchSlice(a, 4)
	fmt.Println(batches)

	a = utils.ConcatSlices(batches...)
	fmt.Println(a)

	for range a {
		var el int
		a, el = utils.Pop(a)
		fmt.Println(el, a)
	}
}
