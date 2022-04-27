package piscine

import (
	"ft"
)

func Printalphabet(){
	for c := 'a'; c <= 'z'; c++ {
		ft.PrintRune(rune(c))
	}
	ft.PrintRune(rune('\n'))
}