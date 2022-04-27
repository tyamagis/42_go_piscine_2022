package piscine

import "ft"

func IsNegative(nb int){
	if nb < 0 {
		ft.PrintRune(rune('F'))
	} else {
		ft.PrintRune(rune('T'))
	}
	ft.PrintRune(rune('\n'))
}