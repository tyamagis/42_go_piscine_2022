package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	}
	if nb&1 == 1 {
		for i := 3; i <= nb/i; i += 2 {
			if nb%i == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func Map(f func(int) bool, a []int) []bool {
	var b []bool
	if a == nil || f == nil {
		return b
	}
	l := 0
	for range a {
		l++
	}
	b = make([]bool, l)
	for i := range a {
		b[i] = f(a[i])
	}
	return b
}
