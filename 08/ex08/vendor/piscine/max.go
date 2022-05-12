package piscine

func Max(a []int) int {
	max := int((^uint(0)) >> 1)
	min := max + 1
	ret := min
	l := 0
	for range a {
		l++
	}
	if l == 0 {
		return 0
	}
	for _, ia := range a {
		if ret < ia {
			ret = ia
		}
	}
	return ret
}
