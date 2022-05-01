package main

import (
	"fmt"
	p "piscine"
)
func main() {
	elems := []string{"ABC", "def", "GHI", "jkl"}
	fmt.Println(p.Join(elems, "_|_"))
}
