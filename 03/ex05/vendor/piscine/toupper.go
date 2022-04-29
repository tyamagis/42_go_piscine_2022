package piscine

func IsLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func ToUpper(s string) string {
	r := []rune(s)
	for i, rc := range r {
		if IsLower(rc) {
			r[i] = rc - 'a' + 'A'
		}
	}
	return string(r)
}
