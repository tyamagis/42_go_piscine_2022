package piscine

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func isAlpha(r rune) bool {
	return (isLower(r) || isUpper(r))
}

func isDigit(r rune) bool {
	return (r >= '0' && r <= '9')
}

func isAlNum(r rune) bool {
	return (isDigit(r) || isAlpha(r))
}

func isTop(r []rune, i int) bool {
	return (i == 0 || !isAlNum(r[i-1]))
}

func toUpper(r rune) rune {
	return rune(r - 'a' + 'A')
}

func toLower(r rune) rune {
	return rune(r - 'A' + 'a')
}

func Capitalize(s string) string {
	r := []rune(s)
	for i, rc := range r {
		if isTop(r, i) && isLower(rc) {
			r[i] = toUpper(rc)
		}
		if !isTop(r, i) && isUpper(rc){
			r[i] = toLower(rc)
		}
	}
	return string(r)
}
