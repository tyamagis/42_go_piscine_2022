package main

import (
	"ft"
	p "piscine"
)
func main() {
	for i := 1; i < 12; i++ {
		ft.PrintRune(p.NRune("123abc👋😀🤚🌐", i)) 
		ft.PrintRune('\n')
	}
}
