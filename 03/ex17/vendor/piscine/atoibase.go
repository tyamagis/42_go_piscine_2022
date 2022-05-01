package piscine

import "ft"

func printStr(s string) {
	for _, c :=range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func swap(a, b *rune) {
	*a, *b = *b, *a
}

func strRev(s string) string {
	l := strLen(s)
	rev := []rune(s)
	for i := 0; i < l / 2; i++ {
		swap(&rev[i], &rev[l - 1 - i])
	}
	return string(rev)
}

func strLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func isDup(s string) bool {
	for i, ci := range s {
		for j, cj := range s {
			if (i <= j) {
				continue
			} else if ci == cj {
				return true
			}
		}
	}
	return false
}

func isSign(s string) bool {
	for _, c := range s {
		if (c == '+' || c == '-') {
			return true
		}
	}
	return false
}

// base
// - include at least 2 chars
// - no duplicate
// - not include [+] [-]
func isBaseValid(base string) bool {
	if strLen(base) < 2 || isDup(base) || isSign(base) {
		return false
	}
	return true
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
