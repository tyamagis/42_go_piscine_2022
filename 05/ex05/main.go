package main

import (
	"fmt"
	p "piscine"
)

func main(){
	nbr, from, to := "00000000", "01", "0123456789"
	fmt.Println(nbr, from, to)
	fmt.Println(p.ConvertBase(nbr, from, to))
	nbr, from, to = "00000010", "01", "0123456789"
	fmt.Println(nbr, from, to)
	fmt.Println(p.ConvertBase(nbr, from, to))
	nbr, from, to = "00000010", "0123456789", "01"
	fmt.Println(nbr, from, to)
	fmt.Println(p.ConvertBase(nbr, from, to))
	nbr, from, to = "00001000", "01", "0123456789"
	fmt.Println(nbr, from, to)
	fmt.Println(p.ConvertBase(nbr, from, to))
	nbr, from, to = "00001024", "0123456789", "01"
	fmt.Println(nbr, from, to)
	fmt.Println(p.ConvertBase(nbr, from, to))
	nbr, from, to = "7fffffff", "0123456789abcdef", "0123456789"
	fmt.Println(nbr, from, to)
	fmt.Println(p.ConvertBase(nbr, from, to))
}