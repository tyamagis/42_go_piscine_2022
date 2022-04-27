package piscine

func IsDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func BasicAtoi2(s string) int {
	ret := 0
	for _, c := range s {
		if IsDigit(c) == false {
			return (0)
		}
		ret = ret*10 + int(c) - int('0')
	}
	return ret
}
