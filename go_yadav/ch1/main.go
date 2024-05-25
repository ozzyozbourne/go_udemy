package main

import (
	f "fmt"
)

var print = f.Printf

const (
	a = 123
)

func main() {
	var b = 123
	var c float32 = a
	//var d float32 = b error conversion is needed only const are untyped not variables

	print("%d %d %f\n", a, b, c)

	//example of switch to replace if else ladder
	const (
		a  = 20
		bc = 40
	)

	switch {
	case a > bc:
		print("a is greater\n")
	case a < bc:
		print("a is less\n")
	default:
		print("a is equal to b")
	}
	i := 0

	//can omit initilization; conditional; increment
	for {
		i++
		if i == 5 {
			break
		}
	}

	//slice
	var s = []int{1, 2, 3, 4, 5, 6, 7}
	print("%d %d\n", len(s), cap(s))
	var w []int
	m := make([]int, 10, 70)
	print("%v %v\n", w, m)
	print("%d %d %d %d\n", len(w), len(m), cap(w), cap(m))

	var p []int
	print("%t\n", p == nil)

	sd := make([]int, 5, 10)
	print("%t\n", sd == nil)

	var ed map[string]int
	we := make(map[int]int, 5)

	print("%v %v\n", len(ed), len(we))
	print("%t %t\n", ed == nil, we == nil)

}
