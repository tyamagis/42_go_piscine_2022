package main

import (
	"fmt"
	p "piscine"
)

func main(){
	s1 := []int{5, 4, 3, 2, 1, 0}
	s2 := []int{5, 3, 4, 2, 1, 0}
	s3 := []int{1, 0, 2, 4, 5, 3}

	p.SortIntegerTable(s1)
	p.SortIntegerTable(s2)
	p.SortIntegerTable(s3)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}