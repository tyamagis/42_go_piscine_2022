package piscine

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isTop(s string, i int) bool {
	if i == 0 {
		return true
	}
}

func toUpper(r rune) rune {
	return rune(r - 'a' + 'A')
}

func Capitalize(s string) string {
	r := []rune(s)
	for i, c := range s {
		if (isTop(s, i) && isLower(c)) {
			r[i] = ToUpper(c)
		}
	}
}
