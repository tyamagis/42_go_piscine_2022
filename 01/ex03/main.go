package main

import (
	"fmt"
	p "piscine"
)

func main(){
	a := 13
	b := 2
	p.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}