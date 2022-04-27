package piscine

import (
	"ft"
)

func PrintReverseAlphabet(){
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(rune(c))
	}
	ft.PrintRune(rune('\n'))
}