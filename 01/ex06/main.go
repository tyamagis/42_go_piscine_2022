package main

import (
	"fmt"
	p "piscine"
)

func main(){
	a, b := 0, 1
	fmt.Println("before : a ", a, ", b ", b, "\n")
	p.Swap(&a, &b)
	fmt.Println("after  : a ", a, ", b ", b, "\n")
}