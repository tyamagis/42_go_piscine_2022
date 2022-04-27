package main

import (
	"fmt"
	p "piscine"
)

func main(){
	fmt.Println(p.StrRev("Hello World !"))
	fmt.Println(p.StrRev("ハローワールド"))
	fmt.Println(p.StrRev("你好，世界。"))
	fmt.Println(p.StrRev("Helló világ"))
	fmt.Println(p.StrRev("мир приветствй"))
	fmt.Println(p.StrRev("👋😀🤚🌐"))
}