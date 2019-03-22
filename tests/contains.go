package main

import (
	"math"
)

type slice []int

func testNoChans(num int) bool {
	var ints slice
	ints = makeTestSlice(num)
	return ints.containsNoChans(1)
}

func (s slice) containsNoChans(num int) bool {
	for _, v := range s {
		if num == v {
			return true
		}
	}
	return false
}

func testChans(num int) bool {
	min := func(a, b int) int {
		return (int)(math.Min((float64)(a), (float64)(b)))
	}
	var ints slice
	ints = makeTestSlice(num)
	ch := make(chan bool, 1)
	slices := 4
	length := len(ints)
	for i := 0; i < slices; i++ {
		start, end := (i/slices)*length, min((i+1/slices)*length, length)
		testSlice := ints[start:end]
		go testSlice.contains(1, ch)
	}
	b, ok := <-ch
	if ok == false {
		return false
	}
	return b
}

func makeTestSlice(num int) []int {
	ints := make([]int, num)
	ints = append(ints, 1)
	return ints
}

func (s slice) contains(n int, ch chan bool) {
	for _, v := range s {
		if n == v {
			ch <- true
			close(ch)
			break
		}
	}
}
