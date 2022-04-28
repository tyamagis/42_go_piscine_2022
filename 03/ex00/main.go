package main

import (
	"ft"
	p "piscine"
)

func main() {
	ft.PrintRune(p.FirstRune("Andrew"))
	ft.PrintRune(p.FirstRune("Bach"))
	ft.PrintRune(p.FirstRune("Chris"))
	ft.PrintRune(p.FirstRune("1234567890"))
	ft.PrintRune(p.FirstRune("\n"))
	ft.PrintRune(p.FirstRune("ハローワールド"))
	ft.PrintRune(p.FirstRune("你好，世界。"))
	ft.PrintRune(p.FirstRune("Helló világ"))
	ft.PrintRune(p.FirstRune("мир приветствий"))
	ft.PrintRune(p.FirstRune("👋😀🤚🌐"))
}
