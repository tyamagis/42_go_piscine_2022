package main

import (
	"fmt"
	p "piscine"
)

func main(){
	l := p.StrLen("Hello World !\n")
	fmt.Println(l)
	l = p.StrLen("ハローワールド\n")
	fmt.Println(l)
	l = p.StrLen("你好，世界。\n")
	fmt.Println(l)
	l = p.StrLen("Helló világ\n")
	fmt.Println(l)
	l = p.StrLen("мир приветствий\n")
	fmt.Println(l)
	l = p.StrLen("👋😀🤚🌐")
	fmt.Println(l)

}