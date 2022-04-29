package main

import (
	"fmt"
	p "piscine"
)
func main() {
	fmt.Println(p.Compare("Hello", "Hello"))
	fmt.Println(p.Compare("Hello", "Hell"))
	fmt.Println(p.Compare("abc", "bc"))
	fmt.Println(p.Compare("bc", "abc"))
	fmt.Println(p.Compare("abc", "acb"))
	fmt.Println(p.Compare("👋😀🤚🌐", "👋😀🤚🌐"))
	fmt.Println(p.Compare("👋😀🤚🌐", "👋😀👋🌐"))
}
