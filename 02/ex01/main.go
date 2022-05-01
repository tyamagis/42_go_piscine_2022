package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.RecursiveFactorial(0))
	fmt.Println(p.RecursiveFactorial(1))
	fmt.Println(p.RecursiveFactorial(2))
	fmt.Println(p.RecursiveFactorial(3))
	fmt.Println(p.RecursiveFactorial(4))
	fmt.Println(p.RecursiveFactorial(5))
	
	fmt.Println(p.RecursiveFactorial(-1))
	fmt.Println(p.RecursiveFactorial(20))
	fmt.Println(p.RecursiveFactorial(21))
}
