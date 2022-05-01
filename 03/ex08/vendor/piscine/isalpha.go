package piscine

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func isAlphabet(r rune) bool {
	return (isLower(r) || isUpper(r))
}

func isDigit(r rune) bool {
	return (r >= '0' && r <= '9')
}
func IsAlpha(s string) bool {
	for _, c := range s {
		if isDigit(c) == false && isAlphabet(c) == false {
			return false
		}
	}
	return true
}
