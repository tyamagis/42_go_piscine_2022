package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println("[(-2), 0] - [(2), 0] =====")
	for i := -2; i <= 2; i++ {
		fmt.Println(p.MakeRange(i, 0))
	}
	fmt.Println("[0, (-2)] - [0, (4)] =====")
	for i := -2; i < 5; i++ {
		fmt.Println(p.MakeRange(0, i))
	}
}
