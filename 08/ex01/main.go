package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.Median(9, 1, 4, 5, 8))
	fmt.Println(p.Median(2, -3, 1, 5, -10))
	fmt.Println(p.Median(654654, -1841964, 564465, -13216, 6549151))
	fmt.Println(p.Median(-65469, -565451, -5445454, -54544555, -545458))
	fmt.Println(p.Median(0, 0, 0, 0, 0))
}
