package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	ans := 1
	for ans < nb / ans {
		ans++
	}
	if ans * ans == nb {
		return ans
	} else {
		return 0
	}
}
