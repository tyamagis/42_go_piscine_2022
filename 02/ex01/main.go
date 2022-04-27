package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.RecursiveFractorial(0))
	fmt.Println(p.RecursiveFractorial(1))
	fmt.Println(p.RecursiveFractorial(2))
	fmt.Println(p.RecursiveFractorial(3))
	fmt.Println(p.RecursiveFractorial(4))
	fmt.Println(p.RecursiveFractorial(5))
	
	fmt.Println(p.RecursiveFractorial(-1))
	fmt.Println(p.RecursiveFractorial(20))
	fmt.Println(p.RecursiveFractorial(21	))
}
