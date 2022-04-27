package main

import (
	"fmt"
	p "piscine"
)

func main() {
	for i := -5; i < 30; i++ {
		fmt.Println(i, p.IsPrime(i))
	}
}
