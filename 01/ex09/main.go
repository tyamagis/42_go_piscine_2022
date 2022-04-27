package main

import (
	"fmt"
	p "piscine"
)

func main(){
	fmt.Println(p.BasicAtoi2("12345"))
	fmt.Println(p.BasicAtoi2("000012345"))
	fmt.Println(p.BasicAtoi2("123450000"))
	fmt.Println(p.BasicAtoi2("000 000012345"))
	fmt.Println(p.BasicAtoi2("12 345"))
	fmt.Println(p.BasicAtoi2("Hello world"))
	fmt.Println(p.BasicAtoi2("+12345"))
	fmt.Println(p.BasicAtoi2("-12345"))
}