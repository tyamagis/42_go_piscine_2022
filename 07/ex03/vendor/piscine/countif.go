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

func CountIf(f func(string) bool, tab []string) int {
	ct := 0
	if f == nil || tab == nil {
		return ct
	}
	for i := range tab {
		if f(tab[i]) == true {
			ct++
		}
	}
	return ct
}
