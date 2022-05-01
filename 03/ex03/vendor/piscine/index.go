package piscine

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func RuneCmp(s1, s2 []rune, i int) bool {
	for j, r := range s2 {
		if s1[i + j] != r {
			return false
		}
	}
	return true
}

func Index(hs, ndl string) int {
	r_hs := []rune(hs)
	len_hs := StrLen(hs)
	r_ndl := []rune(ndl)
	len_ndl := StrLen(ndl)
	if len_hs < len_ndl {
		return -1
	}
	for i := 0; i <= len_hs - len_ndl; i++ {
		if RuneCmp(r_hs, r_ndl, i) {
			return i
		}
	}
	return -1
}
