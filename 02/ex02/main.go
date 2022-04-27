package main

import (
	"fmt"
	p "piscine"
)
func main() {
	fmt.Println(p.IterativePower(0, 0))
	fmt.Println(p.IterativePower(0, 1))
	fmt.Println(p.IterativePower(1, 0))
	fmt.Println(p.IterativePower(1, 1))
	fmt.Println(p.IterativePower(2, 0))
	fmt.Println(p.IterativePower(2, 1))
	fmt.Println(p.IterativePower(2, 2))
	fmt.Println(p.IterativePower(2, 3))
	fmt.Println(p.IterativePower(2, 4))
	fmt.Println(p.IterativePower(10, 3))
	fmt.Println(p.IterativePower(16, 8))
}
