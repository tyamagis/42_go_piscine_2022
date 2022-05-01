package main

import (
	"fmt"
	p "piscine"
)

func main() {
	for i := -3; i <= 16; i++ {
		fmt.Println("sqrt(", i, ") =", p.Sqrt(i))
	}
}
