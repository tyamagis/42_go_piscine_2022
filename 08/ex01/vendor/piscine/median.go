package piscine

func sort(a []int) {
	for i, _ := range a {
		for j, _ := range a {
			if i > j && a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			} else if i < j && a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func Median(a, b, c, d, e int) int {
	ar := []int{a, b, c, d, e}
	sort(ar)
	return ar[2]
}
