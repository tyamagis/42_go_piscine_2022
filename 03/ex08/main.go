package main

import (
	"fmt"
	p "piscine"
)
func main() {
	fmt.Println(p.IsAlpha("IsAlphaFunctionReturnsBool"))
	fmt.Println(p.IsAlpha("It4110wsD1g1ts"))
	fmt.Println(p.IsAlpha("Space isn't allowed."))
	fmt.Println(p.IsAlpha("SymbolAlsoReturnsFalse!!"))
}
