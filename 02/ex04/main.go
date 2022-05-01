package main

import (
	"fmt"
	p "piscine"
)

func main() {
	for i := 0; i <= 30; i++ {
		fmt.Println(p.Fibonacci(i))
	}
}
