package piscine

func isUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func IsUpper(s string) bool {
	for _, c := range s {
		if isUpper(c) == false {
			return false
		}
	}
	return true
}
