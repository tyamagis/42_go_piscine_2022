package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func NRune(s string, n int) rune {
	r := []rune(s)
	l := StrLen(s)
	if n <= 0 || n > l {
		return rune(0)
	} else {
		return r[n - 1]
	}
}
