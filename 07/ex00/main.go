package main

import (
	p "piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	p.ForEach(p.PrintNbr, a)

	b := []int{-9223372036854775808, -1, 0, -1, 9223372036854775807}
	p.ForEach(p.PrintNbr, b)

	// nil slice
	var c []int
	p.ForEach(p.PrintNbr, c)

	// empty slice
	d := make([]int, 0)
	p.ForEach(p.PrintNbr, d)

	e := []int{1, 2, 3, 4, 5, 6}
	var f func(int)
	p.ForEach(f, e)
}
