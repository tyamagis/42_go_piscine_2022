package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 3, 4, 5}
	a3 := []int{0, 1, 1, 2, 3, 3, 4, 5}
	a4 := []int{-100, -99, 99, 100}
	a5 := []int{0, 1, 0, 1, 2, 3}
	a6 := []int{0, 1, 2, -1, 4}
	a7 := []int{0}
	a8 := []int{}
	fmt.Println(p.IsSorted(p.NCompare, a1))
	fmt.Println(p.IsSorted(p.NCompare, a2))
	fmt.Println(p.IsSorted(p.NCompare, a3))
	fmt.Println(p.IsSorted(p.NCompare, a4))
	fmt.Println(p.IsSorted(p.NCompare, a5))
	fmt.Println(p.IsSorted(p.NCompare, a6))
	fmt.Println(p.IsSorted(p.NCompare, a7))
	fmt.Println(p.IsSorted(p.NCompare, a8))
 }
