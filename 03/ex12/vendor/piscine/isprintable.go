package piscine

func isPrintable(r rune) bool {
	return (r >= ' ' && r <= '~')
}

func IsPrintable(s string) bool {
	for _, c := range s {
		if isPrintable(c) == false {
			return false
		}
	}
	return true
}
