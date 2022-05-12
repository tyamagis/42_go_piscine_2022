package piscine

import (
	"ft"
	"os"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func DisplayFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		PrintStr(err.Error())
		return
	}
	PrintStr(string(data))
}
