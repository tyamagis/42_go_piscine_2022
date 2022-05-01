package piscine

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func IsLower(s string) bool {
	for _, c := range s {
		if isLower(c) == false {
			return false
		}
	}
	return true
}
