package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func LastRune(s string) rune {
	r := []rune(s)
	l := StrLen(s)
	return r[l - 1]
}
