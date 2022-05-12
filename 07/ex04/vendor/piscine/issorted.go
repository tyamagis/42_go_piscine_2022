package piscine

func NCompare(a, b int) int {
	return a - b
}

func Length(a []int) (l int) {
	for range a {
		l++
	}
	return l
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if f == nil || a == nil || Length(a) == 0 {
		return false
	}
	for i := range a {
		if i == 0 {
			continue
		}
		if f(a[i-1], a[i]) != -1 {
			return false
		}
	}
	return true
}
