package piscine

func swap(a, b *rune) {
	*a, *b = *b, *a
}

func strRev(s string) string {
	l := strLen(s)
	rev := []rune(s)
	for i := 0; i < l/2; i++ {
		swap(&rev[i], &rev[l-1-i])
	}
	return string(rev)
}

func strLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func isDup(s string) bool {
	for i, ci := range s {
		for j, cj := range s {
			if i <= j {
				continue
			} else if ci == cj {
				return true
			}
		}
	}
	return false
}

func isSign(s string) bool {
	for _, c := range s {
		if c == '+' || c == '-' {
			return true
		}
	}
	return false
}

// base
// - include at least 2 chars
// - no duplicate
// - not include [+] [-]
func isBaseValid(base string) bool {
	if strLen(base) < 2 || isDup(base) || isSign(base) {
		return false
	}
	return true
}

func ItoaBase(nb int, base string) string {
	rb := []rune(base)
	if isBaseValid(base) == false || nb == 0 {
		return "0"
	}
	base_n := strLen(base)
	str := ""
	sign := 1
	if nb < 0 {
		sign = -1
	}
	for nb != 0 {
		mod := nb % base_n
		if mod < 0 {
			mod *= -1
		}
		str += string(rb[mod])
		nb /= base_n
	}
	if sign == -1 {
		str += string('-')
	}
	return strRev(str)
}
