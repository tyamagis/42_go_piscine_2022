package piscine

import (
	"ft"
)

func printStr(s string){
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func RevParams(args []string){
	argc := 0
	for range args {
		argc++
	}
	for argc > 1 {
		printStr(args[argc - 1])
		argc--
	}
}