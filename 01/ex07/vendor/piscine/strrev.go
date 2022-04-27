package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func Swap(a, b *rune) {
	*a, *b = *b, *a
}

func StrRev(s string) string {
	l := StrLen(s)
	rev := []rune(s)
	for i := 0; i < l / 2; i++ {
		Swap(&rev[i], &rev[l - 1 - i])
	}
	return string(rev)
}
