package piscine

func isDigit(r rune) bool {
	return (r >= '0' && r <= '9')
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if isDigit(c) == false {
			return false
		}
	}
	return true
}

func Any(f func(string) bool,  a []string) bool {
	if f == nil || a == nil {
		return false
	}
	for _, s := range a {
		if f(s) == true {
			return true
		}
	}
	return false
}