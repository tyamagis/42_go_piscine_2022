package main

import (
	"fmt"
	p "piscine"
)

func main() {
	for i := -1; i < 600000000000000000; i += (i + 2) {
		fmt.Println(p.FindNextPrime(i))
	}
}
