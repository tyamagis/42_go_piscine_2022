package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 3, 1, 4, 5}
	a3 := []int{}
	fmt.Println(p.IsSorted(p.NCompare, a1))
	fmt.Println(p.IsSorted(p.NCompare, a2))
	fmt.Println(p.IsSorted(p.NCompare, a3))
}
