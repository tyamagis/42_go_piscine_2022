package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.IsNumeric("0012345600"))
	fmt.Println(p.IsNumeric("0.02"))
	fmt.Println(p.IsNumeric("+123"))
	fmt.Println(p.IsNumeric("-123"))
	fmt.Println(p.IsNumeric("7/5"))
}
