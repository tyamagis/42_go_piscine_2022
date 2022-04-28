package main

import (
	"fmt"
	p "piscine"
)
func main() {
	fmt.Println(p.Index("Hello!", "H"))
	fmt.Println(p.Index("Hello!", "e"))
	fmt.Println(p.Index("Hello!", "l"))
	fmt.Println(p.Index("Hello!", "ll"))
	fmt.Println(p.Index("Hello!", "ll!"))
	fmt.Println(p.Index("Hello!", "He1"))
	fmt.Println(p.Index("Hello!", "lo?"))
	fmt.Println(p.Index("Hello!", "oo"))
	fmt.Println(p.Index("👋😀🤚🌐", "👋"))
	fmt.Println(p.Index("👋😀🤚🌐", "😀"))
	fmt.Println(p.Index("👋😀🤚🌐", "🤚"))
	fmt.Println(p.Index("👋😀🤚🌐", "🌐"))
	fmt.Println(p.Index("👋😀🤚🌐", "😀🤚"))
	fmt.Println(p.Index("👋😀🤚🌐", "😀🌐"))
}
