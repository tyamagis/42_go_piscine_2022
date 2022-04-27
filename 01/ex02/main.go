package main

import (
	"fmt"
	p "piscine"
)

func main() {
	a := 13
	b := 2
	var div, mod int
	p.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
