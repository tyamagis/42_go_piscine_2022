package main

import (
	"fmt"
	p "piscine"
)
func main() {
	fmt.Println(p.IsPrintable(" ,	, += --"))
	fmt.Println(p.IsPrintable(" \n"))
	fmt.Println(p.IsPrintable(" "))
}
