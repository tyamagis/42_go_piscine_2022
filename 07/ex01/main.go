package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 17, 19, 23, 29, 30, 31, 32}
	res := p.Map(p.IsPrime, a)
	fmt.Println(res)
}
