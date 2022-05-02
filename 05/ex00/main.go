package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println("[(6), 8] - [(10), 8] =====")
	for i := 6; i <= 10; i++ {
		fmt.Println(p.AppendRange(i, 8))
	}
	fmt.Println("[0, (-2)] - [0, (4)] =====")
	for i := -2; i < 5; i++ {
		fmt.Println(p.AppendRange(0, i))
	}
}
