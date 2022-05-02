package piscine

import "ft"

func printStr(s string) {
	for _, c :=range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func isBaseChar(s, base string) bool {
	for _, c := range s {
		exist := 0
		for _, bc := range base {
			if c == bc {
				exist = 1
			}
		}
		if exist == 0 {
			return false
		}
	}
	return true
}

func isNbrValid(s, base string) bool {
	if strLen(s) == 0 || isSign(s) || isBaseChar(s, base) == false {
		return false
	}
	return true
}

func getBaseIdx(r rune, s string) int {
	idx := 0
	for _, c := range s {
		if (c == r) {
			break
		}
		idx++
	}
	return idx
}

func AtoiBase(s, base string) int {
	if isBaseValid(base) == false || isNbrValid(s, base) == false {
		return 0
	}
	base_n := strLen(base)
	rs := []rune(s)
	ans := 0
	for _, c := range rs {
		ans = ans * base_n + getBaseIdx(c, base)
	}
	return ans
}
