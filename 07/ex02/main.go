package main

import (
	"fmt"
	p "piscine"
)

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"H3110", "h0w", "ar3", "0"}
	tab3 := []string{"Hello", "123", "456", "78q"}
	tab4 := []string{"Hello", "123", "456", "789"}
	tab5 := []string{"", "", "", ""}
	a1 := p.Any(p.IsNumeric, tab1)
	a2 := p.Any(p.IsNumeric, tab2)
	a3 := p.Any(p.IsNumeric, tab3)
	a4 := p.Any(p.IsNumeric, tab4)
	a5 := p.Any(p.IsNumeric, tab5)
	fmt.Println(a1, a2, a3, a4, a5)
}
