package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	ans := 1
	for ans <= nb/ans {
		ans++
	}
	if (ans-1)*(ans-1) == nb {
		return ans-1
	}
	return ans
}
