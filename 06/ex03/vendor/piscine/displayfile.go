package piscine

import (
	"ft"
	"os"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func PrintErr(s string) {
	PrintStr("ERROR: ")
	PrintStr(s)
	PrintStr("\n")
}

func DisplayFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		PrintErr(err.Error())
		return
	}
	PrintStr(string(data))
}
