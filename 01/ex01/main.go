package main

import (
	"fmt"
	p "piscine"
)

func main(){
	a := 0	
	b := &a
	n := &b
	p.UltimatePointOne(&n)
	fmt.Println(a)
}