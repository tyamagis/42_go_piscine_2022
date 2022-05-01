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

func PrintParams(args []string){
	argc := 0
	for range args {
		argc++
	}
	for i := 1; i < argc; i++ {
		printStr(args[i])
	}
}