package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a := []int{654, 456, 546, 645, 465, 564}
	fmt.Println(p.Max(a))

	b := []int{-9223372036854775808}
	fmt.Println(p.Max(b))

	c := []int{}
	fmt.Println(p.Max(c))
}
