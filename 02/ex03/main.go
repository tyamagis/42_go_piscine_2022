package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.RecursivePower(0, 0))
	fmt.Println(p.RecursivePower(0, 1))
	fmt.Println(p.RecursivePower(1, 0))
	fmt.Println(p.RecursivePower(1, 1))
	fmt.Println(p.RecursivePower(2, -1))
	fmt.Println(p.RecursivePower(2, 0))
	fmt.Println(p.RecursivePower(2, 1))
	fmt.Println(p.RecursivePower(2, 2))
	fmt.Println(p.RecursivePower(2, 3))
	fmt.Println(p.RecursivePower(10, 3))
	fmt.Println(p.RecursivePower(16, 8))
}
