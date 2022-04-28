package main

import (
	"ft"
	p "piscine"
)
func main() {
	ft.PrintRune(p.LastRune("Hello!"))
	ft.PrintRune(p.LastRune("world."))
	ft.PrintRune(p.LastRune("ABCDEFG"))
	ft.PrintRune(p.LastRune("Hello"))
	ft.PrintRune(p.LastRune("Hello👋"))
	ft.PrintRune('\n')
}
