package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := p.Unmatch(a)
	fmt.Println(unmatch)
}
