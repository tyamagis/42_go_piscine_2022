package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a := make([]string, 6)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("Size after compacting:", p.Compact(&a))
	for _, v := range a {
		fmt.Println(v)
	}
}
