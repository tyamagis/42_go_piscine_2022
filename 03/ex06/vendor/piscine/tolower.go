package piscine

func IsUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func ToLower(s string) string {
	r := []rune(s)
	for i, rc := range r {
		if IsUpper(rc) {
			r[i] = rc - 'A' + 'a'
		}
	}
	return string(r)
}
