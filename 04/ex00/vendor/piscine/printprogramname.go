package piscine

import (
	"ft"
)

func findName(s string) int {
	idx := 0
	for i, c := range s {
		if c == '/' {
			idx = i
		}
	}
	return idx
}

func strLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func PrintProgramName(s string){
	r := []rune(s)
	idx := findName(s)
	if idx != 0 {
		idx++
	}
	len := strLen(s)

	for i := idx; i < len; i++ {
		ft.PrintRune(r[i])
	}
	ft.PrintRune('\n')
}