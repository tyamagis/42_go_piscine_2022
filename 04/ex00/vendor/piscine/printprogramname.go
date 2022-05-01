package piscine

import (
	"ft"
)

func PrintProgramName(s string){
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}