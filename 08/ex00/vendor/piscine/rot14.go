package piscine

func isAlpha(r rune) bool {
	return ((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'))
}

func Replace(r rune) rune {
	if (r >= 'a' && r <= 'l') || (r >= 'A' && r <= 'L'){
		return r + 14
	} else {
		return r - 12
	}
}

func Rot14(s string) string {
	rs := []rune(s)
	for i, r := range rs {
		if isAlpha(r) {
			rs[i] = Replace(r)
		}
	}
	return string(rs)
}
